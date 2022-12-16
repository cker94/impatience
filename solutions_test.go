package main

import (
	"container/heap"
	"testing"
)

func TestMoveHeap(t *testing.T) {
	expected := [7]int{1, 2, 2, 3, 4, 6, 7}
	next := expected[:]
	h := &MoveHeap{
		make([]*Move, 6),
		make([]*Move, 2),
		make([]*Move, 4),
	}
	heap.Init(h)
	heap.Push(h, make([]*Move, 2))
	heap.Push(h, make([]*Move, 3))
	heap.Push(h, make([]*Move, 1))
	heap.Push(h, make([]*Move, 7))

	if h.Len() != len(expected) {
		t.Errorf("Setup error: heap inputs and expected list size mismatch: %d inputs; %d expected.", h.Len(), len(expected))
		t.SkipNow()
	}

	results := make([]int, 0, h.Len())
	for h.Len() > 0 {
		result := len(heap.Pop(h).([]*Move))
		results = append(results, result)
		if next[0] != result {
			if !t.Failed() {
				t.Error("Heap order didn't match expected.")
			}
		}
		next = next[1:]
	}
	if t.Failed() {
		t.Log("Results:  ", results)
		t.Log("Expected: ", expected)
	}
}

func TestSortedMoveSets(t *testing.T) {
	expected := [7]int{1, 2, 2, 3, 4, 6, 7}
	next := expected[:]
	s := NewSortedMoveSets(
		make([]*Move, 6),
		make([]*Move, 2),
		make([]*Move, 4),
	)
	s.Push(make([]*Move, 2))
	s.Push(make([]*Move, 3))
	s.Push(make([]*Move, 1))
	s.Push(make([]*Move, 7))

	if s.Len() != len(expected) {
		t.Errorf("Setup error: heap inputs and expected list size mismatch: %d inputs; %d expected.", s.Len(), len(expected))
		t.SkipNow()
	}

	results := make([]int, 0, s.Len())
	for s.Len() > 0 {
		result := len(s.Pop())
		results = append(results, result)
		if next[0] != result {
			if !t.Failed() {
				t.Error("Heap order didn't match expected.")
			}
		}
		next = next[1:]
	}
	if t.Failed() {
		t.Log("Results:  ", results)
		t.Log("Expected: ", expected)
	}
}
