package main

import "testing"

func TestSort(t *testing.T) {
	a := []int{3, 4, 2, 1, 5, 0}

	heap := NewHeap(a)

	heap.Sort()

	for i := 1; i < len(heap.Values); i++ {
		if heap.Values[i] != i-1 {
			t.Error(heap.Values[i], "!= ", i)
		}
	}
}
