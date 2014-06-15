package main

import (
	"os"
)

type BufferHeap struct {
	bp   *BufferPool
	Size int
}

func NewBufferHeap(file string, recordSize, blockSize, capacity int) *BufferHeap {

	info, err := os.Stat(file)

	if err != nil {
		panic(err)
	}

	size := int(info.Size() / int64(recordSize))

	return &BufferHeap{
		NewBufferPool(file, recordSize, blockSize, capacity),
		size,
	}
}

func (heap *BufferHeap) Len() int {
	return heap.Size
}

func (heap *BufferHeap) Less(i, j int) bool {
	i64 := int64(i)
	j64 := int64(j)
	return heap.bp.GetRecord(i64).Key < heap.bp.GetRecord(j64).Key
}

func (heap *BufferHeap) Swap(i, j int) {
	i64 := int64(i)
	j64 := int64(j)
	iRecord := heap.bp.GetRecord(i64)
	jRecord := heap.bp.GetRecord(j64)

	heap.bp.WriteRecord(jRecord, i64)
	heap.bp.WriteRecord(iRecord, j64)
}

func (heap *BufferHeap) Shutdown() {
	heap.bp.Shutdown()
}
