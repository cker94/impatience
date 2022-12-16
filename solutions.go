// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type MoveHeap [][]*Move

func (h MoveHeap) Len() int {
	return len(h)
}
func (h MoveHeap) Less(i, j int) bool {
	return len(h[i]) < len(h[j])
}
func (h MoveHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MoveHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.î
	*h = append(*h, x.([]*Move))
}

func (h *MoveHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// A simplified interface for  MoveHeap.
type SortedMoveSets struct {
	sets MoveHeap
}

func NewSortedMoveSets(m ...[]*Move) *SortedMoveSets {
	s := new(SortedMoveSets)
	s.sets = MoveHeap{}
	if len(m) > 0 {
		s.sets = append(s.sets, m...)
		heap.Init(&s.sets)
	}
	return s
}

func (s SortedMoveSets) Len() int {
	return s.sets.Len()
}

func (s *SortedMoveSets) Push(x []*Move) {
	heap.Push(&s.sets, x)
}

func (s *SortedMoveSets) Pop() []*Move {
	return s.sets.Pop().([]*Move)
}
