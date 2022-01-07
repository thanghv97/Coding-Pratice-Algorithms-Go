package trie

type BasicTrieNode struct {
	child     [26]*BasicTrieNode
	isWordEnd bool
}

type BasicTrie struct {
	root *BasicTrieNode
}

func NewBasicTrie() *BasicTrie {
	return &BasicTrie{
		root: &BasicTrieNode{},
	}
}

func (t *BasicTrie) Insert(s string) {
	node := t.root
	for _, ch := range s {
		index := ch - 'a'
		if node.child[index] == nil {
			node.child[index] = &BasicTrieNode{}
		}
		node = node.child[index]
	}
	node.isWordEnd = true
}

func (t *BasicTrie) Find(s string) bool {
	node := t.root
	for _, ch := range s {
		index := ch - 'a'
		if node.child[index] == nil {
			return false
		}
		node = node.child[index]
	}

	return node.isWordEnd
}
