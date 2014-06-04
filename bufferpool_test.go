package main

import "testing"

func TestGet(t *testing.T) {
	bufferPool := NewBufferPool("./Test100.dat", 4, 4096, 1)
	defer bufferPool.Shutdown()

	rtn := make([]byte, 4)
	bufferPool.Get(&rtn, 0)
	bufferPool.Get(&rtn, 1024)
}
