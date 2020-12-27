package tree

type (
	Trie struct {
		root *trieNode
		size int
		r    int
	}
	trieNode struct {
		val      interface{}
		child    []*trieNode
		childCnt int
	}
)

func NewTrie(R int) *Trie {
	return &Trie{
		nil,
		0,
		R,
	}
}

func newTrieNode(R int) *trieNode {
	return &trieNode{
		nil,
		make([]*trieNode, R),
		0,
	}
}

func (trie *Trie) Get(key string) interface{} {
	node := trie.get(trie.root, []rune(key), 0)
	if node == nil {
		return nil
	}
	return node.val
}

func (trie *Trie) get(parent *trieNode, s []rune, i int) *trieNode {
	if parent == nil {
		return nil
	}
	if i == len(s) {
		return parent
	}
	c := s[i]
	return trie.get(parent.child[c], s, i+1)
}

func (trie *Trie) Delete(key string) {
	trie.root = trie.delete(trie.root, []rune(key), 0)
}

func (trie *Trie) delete(parent *trieNode, key []rune, i int) *trieNode {
	if parent == nil {
		return nil
	}
	if i == len(key) {
		parent.val = nil
	} else {
		c := key[i]
		exist := parent.child[c] != nil
		if exist {
			parent.child[c] = trie.delete(parent.child[c], key, i+1)
			if parent.child[c] == nil {
				parent.childCnt--
			}
		}
	}

	//	remove unused node
	if parent.val == nil && parent.childCnt == 0 {
		return nil
	}
	return parent
}

func (trie *Trie) LongestPrefixOf(key string) string {
	runes := []rune(key)
	length := trie.longestPrefixOf(trie.root, runes, 0)
	return string(runes[0:length])
}

func (trie *Trie) longestPrefixOf(parent *trieNode, key []rune, i int) int {
	if parent == nil {
		return 0
	}
	if i == len(key) {
		return 0
	}
	c := key[i]
	if parent.child[c] == nil {
		return 0
	}
	return 1 + trie.longestPrefixOf(parent.child[c], key, i+1)
}

func (trie *Trie) Put(key string, val interface{}) {
	trie.root = trie.put(trie.root, []rune(key), val, 0)
}

func (trie *Trie) put(parent *trieNode, key []rune, val interface{}, i int) *trieNode {
	if parent == nil {
		parent = newTrieNode(trie.r)
	}
	if i == len(key) {
		parent.val = val
		return parent
	}
	if parent.child[key[i]] == nil {
		parent.childCnt++
	}
	parent.child[key[i]] = trie.put(parent.child[key[i]], key, val, i+1)
	return parent
}
