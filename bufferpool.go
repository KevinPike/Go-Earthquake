package main

import "os"

const (
	KeySize   = 2
	DataSize  = 2
	TotalSize = KeySize + DataSize
	BlockSize = 4096
)

type BufferPool struct {
	File       *os.File
	Blocks     []Block
	FreeBlocks int
}

type Block struct {
	Bytes  []byte
	Offset int64
	Used   int64
}

func NewBufferPool(path string, size int) *BufferPool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	blocks := make([]Block, size)

	return &BufferPool{file, blocks, size}
}

func (b *BufferPool) Get(index int64) {
	var offset int64 = index * TotalSize

	for i := 0; i < len(b.Blocks); i++ {
		block := b.Blocks[i]

		diff := block.Offset - offset
		if diff > -1024 && diff < 1 {

		}
	}
}

func (b *BufferPool) Load(offset int64) {
	bytes := make([]byte, BlockSize)
	_, err := b.File.ReadAt(bytes, offset)
	if err != nil {
		panic(err)
	}

	block := NewBlock(bytes, offset)
	if b.FreeBlocks > 0 {
		// If there are blocks left, insert!
		index := len(b.Blocks) - b.FreeBlocks
		b.Blocks[index] = *block
	} else {
		leastIndex := 0
		var least int64 = 9223372036854775807
		for i := 0; i < len(b.Blocks); i++ {
			used := b.Blocks[i].Used
			if used < least {
				least = used
				leastIndex = i
			}
		}
		b.Blocks[leastIndex] = *block
	}
}

func NewBlock(bytes []byte, offset int64) *Block {
	return &Block{bytes, offset, 0}
}
