package tree

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie(256)
	trie.Put("123", 123)
	trie.Put("12", 12)
	trie.Put("ab", "ab")
	v := trie.Get("12")
	assertTrue(v == 12, t)
	v = trie.Get("123")
	assertTrue(v == 123, t)
	key := trie.LongestPrefixOf("1234")
	assertTrue(key == "123", t)
	trie.Delete("123")
	v = trie.Get("123")
	assertTrue(v == nil, t)
	v = trie.Get("12")
	assertTrue(v == 12, t)
	key = trie.LongestPrefixOf("1234")
	assertTrue(key == "12", t)
	key = trie.LongestPrefixOf("ab")
	assertTrue(key == "ab", t)
}

func assertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
