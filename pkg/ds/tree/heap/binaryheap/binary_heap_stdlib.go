// package binaryheap
package binaryheap

import "fmt"

/************************************************************
HEAP BY STANDARD LIB OF GO: container/heap

************************************************************/
type MinHeapGoLib []int

// Add necessary function for standard lib of Go
func (h MinHeapGoLib) Len() int           { return len(h) }
func (h MinHeapGoLib) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapGoLib) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapGoLib) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapGoLib) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeapGoLib) Log() {
	fmt.Print("Heap: ")
	for i := 0; i < h.Len(); i++ {
		fmt.Printf(" %d", (*h)[i])
	}
	fmt.Printf(" => Min: %d\n", (*h)[0])
}
