package hash

type HashTable struct {
	size int
	data map[int][][2]string // [][2]string to store key-value pairs
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size: size,
		data: make(map[int][][2]string, size),
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = hash + int(char)
	}
	return hash % ht.size
}

func (ht *HashTable) Set(key, value string) {
	index := ht.hash(key)

	for i, pair := range ht.data[index] {
		if pair[0] == key {
			ht.data[index][i][1] = value // Update existing key
			return
		}
	}

	if _, ok := ht.data[index]; !ok {
		ht.data[index] = make([][2]string, 0) // Initialize if not present
	}

	ht.data[index] = append(ht.data[index], [2]string{key, value}) // Add new key-value pair
}

func (ht *HashTable) Get(key string) (string, bool) {
	index := ht.hash(key)

	if pairs, ok := ht.data[index]; ok {
		for _, pair := range pairs {
			if pair[0] == key {
				return pair[1], true // Return value if key found
			}
		}
	}

	return "", false // Return empty string and false if key not found
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)

	if pairs, ok := ht.data[index]; ok {
		for i, pair := range pairs {
			if pair[0] == key {
				ht.data[index] = append(pairs[:i], pairs[i+1:]...) // Remove the key-value pair
				return true
			}
		}
	}

	return false // Return false if key not found
}
