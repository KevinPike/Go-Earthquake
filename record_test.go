package main

import (
	"testing"
)

func TestToBytes(t *testing.T) {
	key := []byte{0x18, 0x2d}
	value := []byte{0x44, 0x54}
	record := NewRecord(key, value)
	bytes := record.ToBytes()

	if key[0] != bytes[0] || key[1] != bytes[1] {
		t.Error("Record binary key does not match")
	}
	if value[0] != bytes[2] || value[1] != bytes[3] {
		t.Error("Record binary value does not match")
	}

	record2 := NewRecord(bytes[0:2], bytes[2:4])

	if record.Key != record2.Key || record.Value != record.Value {
		t.Error("Records do not match")
	}
}
