package trees

type Trie struct {
	Value    string
	children map[string]*Trie
}

func NewTrie() *Trie {
	return &Trie{"", make(map[string]*Trie)}
}

func (t *Trie) Insert(word string) {
	t.insert(word, word)
}

func (t *Trie) insert(word, original string) {
	if len(word) == 0 {
		t.Value = original
		return
	}

	first := string(word[0])
	if _, ok := t.children[first]; !ok {
		t.children[first] = NewTrie()
	}
	if len(word) > 0 {
		t.children[first].insert(word[1:], original)
	}
}

func (t *Trie) Search(word string) bool {
	return t.search(word, word)
}

func (t *Trie) search(word, original string) bool {
	if t.Value == original {
		return true
	}

	if len(word) == 0 {
		return false
	}

	first := string(word[0])
	if child, ok := t.children[first]; ok {
		return child.search(word[1:], original)
	}

	return false
}

func (t *Trie) StartsWith(word string) bool {
	if len(word) == 0 {
		return true
	}
	first := string(word[0])
	if child, ok := t.children[first]; ok {
		return child.StartsWith(word[1:])
	}

	return false
}

func (t *Trie) Delete(word string) {
	t.delete(word, word)
}

func (t *Trie) delete(word, original string) bool {
	if len(word) == 0 {
		return true
	}

	first := string(word[0])
	continueDelete := false
	if child, ok := t.children[first]; ok {
		continueDelete = child.delete(word[1:], original) && (child.Value == "" || child.Value == original)
		if continueDelete {
			t.children[first] = nil
		}
	}
	return continueDelete
}
