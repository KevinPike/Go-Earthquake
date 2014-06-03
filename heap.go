package main

type Heap struct {
	Values []int
	Size   int
}

func NewHeap(a []int) *Heap {
	return &Heap{a, len(a) - 1}
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

func (heap Heap) MaxHeapify(i int) {
	l := left(i)
	r := right(i)

	a := heap.Values

	largest := i
	if l <= heap.Size && a[l] > a[i] {
		largest = l
	}

	if r <= heap.Size && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]

		heap.MaxHeapify(largest)
	}
}

func (a Heap) Build() {
	for i := a.Size / 2; i > 0; i-- {
		a.MaxHeapify(i)
	}
}

func (heap Heap) Sort() {
	heap.Build()

	a := heap.Values
	for i := heap.Size; i > 1; i-- {
		a[1], a[i] = a[i], a[1]
		heap.Size--
		heap.MaxHeapify(1)
	}
}
