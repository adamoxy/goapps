package main

import (
	"container/heap"
	"fmt"
)

type CharFrequency struct {
	Char  rune
	Count int
}

// Define a type that implements heap.Interface
type MaxHeap []CharFrequency

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count } // Max-Heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeString(s string) string {
	// Count the frequency of each character
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}

	// Build a max heap based on frequencies
	h := &MaxHeap{}
	heap.Init(h)
	for char, count := range frequency {
		heap.Push(h, CharFrequency{char, count})
	}

	var result []rune
	var prevChar CharFrequency

	// Construct the result string
	for h.Len() > 0 {
		charFreq := heap.Pop(h).(CharFrequency)
		result = append(result, charFreq.Char)
		charFreq.Count--

		if prevChar.Count > 0 {
			heap.Push(h, prevChar)
		}

		prevChar = charFreq

		// Check if the rearrangement is not possible
		if h.Len() == 0 && prevChar.Count > 0 {
			return ""
		}
	}

	return string(result)
}

func main1() {
	fmt.Println(rearrangeString("aab"))  // Output: "aba" or "bab"
	fmt.Println(rearrangeString("adam")) // Output: ""
}
