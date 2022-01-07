package trie

type TreeInterface interface {
	Insert(s string)
	Find(s string) bool
}
