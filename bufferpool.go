package main

import (
	"container/list"
	"os"
)

type Block struct {
	Bytes  []byte
	Offset int64
}

func NewBlock(bytes []byte, offset int64) *Block {
	return &Block{bytes, offset}
}

// Read bytes at offset
func (b *Block) Get(bytes *[]byte, offset int64) {
	numBytes := int64(len(*bytes))
	endOffset := offset + numBytes
	*bytes = b.Bytes[offset:endOffset]
}

func (b *Block) Put(bytes []byte, offset int64) {
	for i := 0; i < len(bytes); i++ {
		b.Bytes[offset+int64(i)] = bytes[i]
	}
}

type BufferPool struct {
	File       *os.File
	Blocks     map[int64]*BlockElement
	List       *list.List
	RecordSize int64
	BlockSize  int64
	Capacity   int
}

type BlockElement struct {
	Block       *Block
	listElement *list.Element
}

func NewBufferPool(path string, recordSize, blockSize, size int) *BufferPool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return &BufferPool{
		File:       file,
		Blocks:     make(map[int64]*BlockElement),
		List:       list.New(),
		RecordSize: int64(recordSize),
		BlockSize:  int64(blockSize),
		Capacity:   size,
	}
}

func (b *BufferPool) Get(rtn *[]byte, index int64) {
	var totalOffset int64 = index * b.RecordSize

	blockOffset := totalOffset / b.BlockSize * b.BlockSize
	indexInBlock := totalOffset % b.BlockSize

	blockElement, exists := b.Blocks[blockOffset]
	if exists == false {
		blockElement = b.Load(blockOffset)
	}

	b.List.MoveToFront(blockElement.listElement)
	block := blockElement.Block

	block.Get(rtn, indexInBlock)
}

/*
Loads a block into the buffer pool.
If no available slots are available, a block is evicted by LRU.
Offset is the byte offset of the block to be loaded.
*/
func (b *BufferPool) Load(offset int64) *BlockElement {
	bytes := make([]byte, b.BlockSize)
	_, err := b.File.ReadAt(bytes, offset)
	if err != nil {
		panic(err)
	}

	// Evict
	if b.List.Len() == b.Capacity {
		tail := b.List.Back()
		blockElement := b.List.Remove(tail).(*BlockElement)
		delete(b.Blocks, blockElement.Block.Offset)
		b.Write(blockElement.Block)
	}

	block := NewBlock(bytes, offset)
	blockElement := &BlockElement{Block: block}
	blockElement.listElement = b.List.PushFront(blockElement)
	b.Blocks[block.Offset] = blockElement
	return blockElement
}

func (b *BufferPool) Write(block *Block) {

}

func (b *BufferPool) IsEmpty() bool {
	return b.List.Len() == 0
}

func (b *BufferPool) Shutdown() {
	// Write all blocks
	//for item := b.List.Front(); b.List.Len() != 0; item = b.List.Front() {
	//	b.Write(item.Value.(*BlockElement).Block)
	//}
	b.File.Close()
}
