package main

import "testing"

func TestSort(t *testing.T) {
	a := []int{0, 3, 2, 4, 1, 5}

	heap := NewHeap(a)

	heap.Sort()

	for i := 1; i < len(heap.Values); i++ {
		if heap.Values[i] != i {
			t.Error("Sort failed")
		}
	}
}
