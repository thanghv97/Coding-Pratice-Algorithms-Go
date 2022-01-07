package binaryheap_test

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/thanghv97/Coding-Pratice-Algorithms-Go/pkg/ds/tree/heap/binaryheap"
)

func TestBinaryHeapStdLib(t *testing.T) {
	// Init Min Heap
	h := &binaryheap.MinHeapGoLib{2, 1, 5, 10}
	heap.Init(h)
	h.Log()

	// Push 3
	fmt.Println("+++ Test Push 3")
	heap.Push(h, 3)
	h.Log()
	assertSame(t, (*h)[0], 1)

	// Pop
	min := heap.Pop(h)
	fmt.Printf("+++ Test Pop => Discard Minimum: %d\n", min)
	h.Log()
	assertSame(t, (*h)[0], 2)

	// Fix
	(*h)[3] = 1
	heap.Fix(h, 3)
	fmt.Println("+++ Test Fix => Value at position 3 changed from 10 to 1")
	h.Log()
	assertSame(t, (*h)[0], 1)

	// Remove
	heap.Remove(h, 1)
	fmt.Println("+++ Test Remove => Discard position 1")
	h.Log()
	assertSame(t, (*h)[0], 1)
}

func assertSame(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("Test Error: expected '%d', got '%d'", a, b)
	}
}
