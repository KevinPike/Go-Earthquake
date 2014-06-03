package main

import (
	"fmt"
	"os"
)

const (
	KeySize   = 2
	DataSize  = 2
	TotalSize = KeySize + DataSize
	BlockSize = 4096
)

type Block struct {
	Bytes  []byte
	Offset int64
	Used   int64
}

func NewBlock(bytes []byte, offset int64) *Block {
	return &Block{bytes, offset, 0}
}

// Read bytes at offset
func (b *Block) Get(bytes *[]byte, offset int64) {
	numBytes := int64(len(*bytes))
	endOffset := offset + numBytes
	*bytes = b.Bytes[offset:endOffset]
}

type BufferPool struct {
	File       *os.File
	Blocks     []Block
	UsedBlocks int
}

func NewBufferPool(path string, size int) *BufferPool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	blocks := make([]Block, size)

	return &BufferPool{file, blocks, 0}
}

func (b *BufferPool) Get(rtn *[]byte, index int64) {
	var totalOffset int64 = index * TotalSize

	found := false
	blockOffset := totalOffset / BlockSize
	indexInBlock := totalOffset % BlockSize
	// Iterate through all blocks,
	// looking for one whos range contains our index
	for i := 0; i < len(b.Blocks) && !b.IsEmpty(); i++ {
		block := b.Blocks[i]

		if block.Offset == blockOffset {
			found = true
			block.Get(rtn, indexInBlock)
			break
		}
	}

	// If not found, load block into buffer pool
	if !found {
		fmt.Println("Loading block for offset ", totalOffset)
		block := b.Load(blockOffset)
		block.Get(rtn, indexInBlock)
	}
}

/*
Loads a block into the buffer pool.
If no available slots are available, a block is evicted by LRU.
Offset is the byte offset of the block to be loaded.
*/
func (b *BufferPool) Load(offset int64) *Block {
	bytes := make([]byte, BlockSize)
	_, err := b.File.ReadAt(bytes, offset)
	if err != nil {
		panic(err)
	}

	block := NewBlock(bytes, offset)
	if b.UsedBlocks < len(b.Blocks) {
		// If there are blocks left, insert!
		b.Blocks[b.UsedBlocks] = *block
		b.UsedBlocks = b.UsedBlocks + 1
	} else {
		// Eviction: LRU
		leastIndex := 0
		var least int64 = 9223372036854775807
		for i := 0; i < len(b.Blocks); i++ {
			used := b.Blocks[i].Used
			if used < least {
				least = used
				leastIndex = i
			}
		}
		// Need to load existing block back into memory
		b.Blocks[leastIndex] = *block
	}

	return block
}

func (b *BufferPool) IsEmpty() bool {
	return b.UsedBlocks == 0
}

func (b *BufferPool) Shutdown() {
	b.File.Close()
}
