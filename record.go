package main

import (
	"bytes"
	"encoding/binary"
)

const (
	KEY_SIZE   = 2
	VALUE_SIZE = 2
)

type Record struct {
	Key   int16
	Value int16
}

func NewRecord(key []byte, value []byte) *Record {
	var k, v int16
	reader := bytes.NewReader(key)
	err := binary.Read(reader, binary.BigEndian, &k)
	if err != nil {
		panic(err)
	}

	reader = bytes.NewReader(value)
	err = binary.Read(reader, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}

	record := &Record{k, v}
	return record
}

func (record Record) ToBytes() []byte {
	w := new(bytes.Buffer)
	err := binary.Write(w, binary.BigEndian, record.Key)
	if err != nil {
		panic(err)
	}

	err = binary.Write(w, binary.BigEndian, record.Value)
	if err != nil {
		panic(err)
	}

	return w.Bytes()
}
