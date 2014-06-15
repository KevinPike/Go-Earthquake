package main

type Block struct {
	Bytes  []byte
	Offset int64 //number of bytes from beginning of file
}

func NewBlock(bytes []byte, offset int64) *Block {
	return &Block{bytes, offset}
}

// Read len([]bytes) bytes from block at index within block
func (b *Block) Get(bytes *[]byte, index int64) {
	numBytes := int64(len(*bytes))
	endOffset := index + numBytes
	*bytes = b.Bytes[index:endOffset]
}

// Updates b.Bytes with bytes at index within block
func (b *Block) Put(bytes []byte, index int64) {
	for i := 0; i < len(bytes); i++ {
		b.Bytes[index+int64(i)] = bytes[i]
	}
}
