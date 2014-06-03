package main

import (
	"fmt"
)

func main() {
	a := []int{0, 2, 1, 3, 4, 6, 5, 7}
	heap := NewHeap(a)
	fmt.Println(heap.Values)
	heap.Build()
	fmt.Println(heap.Values)
	heap.Sort()
	fmt.Println(heap.Values)
}
