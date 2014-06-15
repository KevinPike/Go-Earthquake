package main

import "testing"

func TestGet(t *testing.T) {
	bufferPool := NewBufferPool("./Test100.dat", 4, 4096, 1)
	defer bufferPool.Shutdown()

	bufferPool.GetRecord(0)
	bufferPool.GetRecord(1024)
}
