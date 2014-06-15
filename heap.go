package main

type Heap struct {
	Values []int
}

func (heap *Heap) Len() int {
	return len(heap.Values)
}

func (heap *Heap) Less(i, j int) bool {
	return heap.Values[i] < heap.Values[j]
}

func (heap *Heap) Swap(i, j int) {
	heap.Values[i], heap.Values[j] = heap.Values[j], heap.Values[i]
}

func NewHeap(a []int) *Heap {
	return &Heap{a}
}
