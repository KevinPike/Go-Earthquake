package main

import (
	"sort"
)

type Heap struct {
	Values []int
	Size   int
}

func (heap Heap) Len() int {
	return heap.Size
}

func (heap Heap) Less(i, j int) bool {
	return heap.Values[i] < heap.Values[j]
}

func (heap Heap) Swap(i, j int) {
	heap.Values[i], heap.Values[j] = heap.Values[j], heap.Values[i]
}

func NewHeap(a []int) *Heap {
	b := []int{0}
	b = append(b, a...)
	return &Heap{b, len(b) - 1}
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func MaxHeapify(sortInterface sort.Interface, i int) {
	l := left(i)
	r := right(i)

	largest := i
	if l <= sortInterface.Len() && sortInterface.Less(i, l) {
		largest = l
	}

	if r <= sortInterface.Len() && sortInterface.Less(largest, r) {
		largest = r
	}

	if largest != i {
		sortInterface.Swap(i, largest)

		MaxHeapify(sortInterface, largest)
	}
}

func Build(sortInterface sort.Interface) {
	for i := sortInterface.Len() / 2; i > 0; i-- {
		MaxHeapify(sortInterface, i)
	}
}

func (heap Heap) Sort() {
	Build(heap)

	for i := heap.Len(); i > 1; i-- {
		heap.Swap(1, i)
		heap.Size--
		MaxHeapify(heap, 1)
	}
}
