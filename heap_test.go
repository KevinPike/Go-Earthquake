package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{3, 4, 2, 1, 5, 0}

	heap := NewHeap(a)

	HeapSort(heap)

	for i := 0; i < len(heap.Values); i++ {
		if heap.Values[i] != i {
			t.Error(heap.Values[i], "!= ", i)
		}
	}
}
