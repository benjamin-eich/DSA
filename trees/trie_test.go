package trees

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")

	if trie.Search("apple") != true {
		t.Error("trie should contain apple but Search() returned false")
	}

	if trie.Search("app") == true {
		t.Error("trie should not contain app but Search() returned true")
	}

	if trie.StartsWith("app") != true {
		t.Error("trie should contain entries that start with app but StartsWith() returned false")
	}

	if trie.StartsWith("big") != false {
		t.Error("trie should not contain entries that start with big but StartsWith() returned true")
	}

	trie.Insert("app")

	if trie.Search("app") != true {
		t.Error("trie should contain app but Search() returned false")
	}
}
