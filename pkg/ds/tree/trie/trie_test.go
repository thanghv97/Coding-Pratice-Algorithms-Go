package trie_test

import (
	"fmt"
	"testing"

	"github.com/thanghv97/Coding-Pratice-Algorithms-Go/pkg/ds/tree/trie"
	tv "github.com/thanghv97/Coding-Pratice-Algorithms-Go/pkg/ds/tree/trie/version"
)

func TestBasicTree(t *testing.T) {
	fmt.Println("=== Init Basic Trie ===")
	runTest(t, tv.NewBasicTrie())
}

func runTest(t *testing.T, trie trie.TreeInterface) {
	// Test Add
	records := []string{"abc", "test", "abc", "testabc", "abcd"}
	fmt.Printf("=== Add records{%v} to trie\n", records)
	for _, s := range records {
		trie.Insert(s)
	}

	// Test Find
	if trie.Find("abc") {
		fmt.Println("=== Find 'abc' => TRUE ")
	} else {
		fmt.Println("=== Find 'abc' => FALSE ")
	}

	if trie.Find("a") {
		fmt.Println("=== Find 'a' => TRUE ")
	} else {
		fmt.Println("=== Find 'a' => FALSE ")
	}

	if trie.Find("abcd") {
		fmt.Println("=== Find 'abcd' => TRUE ")
	} else {
		fmt.Println("=== Find 'abcd' => FALSE ")
	}
}
