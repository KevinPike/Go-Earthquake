package main

func main() {
	path := "./Test100 Copy.dat"

	bufferHeap := NewBufferHeap(path, 4, 4096, 20)
	defer bufferHeap.Shutdown()

	HeapSort(bufferHeap)
}
