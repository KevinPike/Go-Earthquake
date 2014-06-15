package main

import (
	"testing"
)

const (
	TestFile = "./TestFile.dat"
)

func TestGet(t *testing.T) {
	bufferPool := NewBufferPool(TestFile, 4, 4096, 1)
	defer bufferPool.Shutdown()

	bufferPool.GetRecord(0)
	bufferPool.GetRecord(1023)
}

func TestWrite(t *testing.T) {
	bufferPool := NewBufferPool(TestFile, 4, 4096, 1)
	defer bufferPool.Shutdown()

	a := bufferPool.GetRecord(0)
	b := bufferPool.GetRecord(1023)

	// Write a to b and b to a
	bufferPool.WriteRecord(a, 1023)
	bufferPool.WriteRecord(b, 0)

	newA := bufferPool.GetRecord(0)
	newB := bufferPool.GetRecord(1023)

	if !newA.Equals(b) || !newB.Equals(a) {
		t.Error("Write did not succeed")
	}
}

func TestFlush(t *testing.T) {
	bufferPool := NewBufferPool(TestFile, 4, 4096, 1)

	a := bufferPool.GetRecord(0)
	b := bufferPool.GetRecord(1023)

	// Write a to b and b to a
	bufferPool.WriteRecord(a, 1023)
	bufferPool.WriteRecord(b, 0)

	bufferPool.Shutdown()

	bufferPool = NewBufferPool(TestFile, 4, 4096, 1)
	defer bufferPool.Shutdown()

	newA := bufferPool.GetRecord(0)
	newB := bufferPool.GetRecord(1023)

	if !newA.Equals(b) || !newB.Equals(a) {
		t.Error("Flush did not succeed")
	}
}
