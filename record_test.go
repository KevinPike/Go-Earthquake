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

func TestEquals(t *testing.T) {
	key := []byte{0x18, 0x2d}
	value := []byte{0x44, 0x54}
	record1 := NewRecord(key, value)
	record2 := NewRecord(key, value)

	if !record1.Equals(record2) {
		t.Error("Equals failed")
	}
}

func TestNotEquals(t *testing.T) {
	key := []byte{0x18, 0x2d}
	value := []byte{0x44, 0x54}
	key2 := []byte{0x18, 0x20}
	value2 := []byte{0x44, 0x40}
	record1 := NewRecord(key, value)
	record2 := NewRecord(key2, value2)

	if record1.Equals(record2) {
		t.Error("Equals failed")
	}
}
