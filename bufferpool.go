package main

import (
	"container/list"
	"os"
)

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

func (b *BufferPool) GetRecord(index int64) *[]byte {
	block, indexInBlock := b.getBlock(index)

	rtn := make([]byte, b.RecordSize)
	block.Get(&rtn, indexInBlock)
	return &rtn
}

func (b *BufferPool) WriteRecord(rtn *[]byte, index int64) {
	block, indexInBlock := b.getBlock(index)

	block.Put(*rtn, indexInBlock)
}

func (b *BufferPool) getBlock(recordIndex int64) (*Block, int64) {
	var totalOffsetInFile int64 = recordIndex * b.RecordSize

	blockOffset := totalOffsetInFile / b.BlockSize * b.BlockSize
	indexInBlock := totalOffsetInFile % b.BlockSize

	// Check if block exists in cache
	blockElement, exists := b.Blocks[blockOffset]
	if exists == false {
		blockElement = b.load(blockOffset)
	}

	b.List.MoveToFront(blockElement.listElement)
	block := blockElement.Block
	return block, indexInBlock
}

/*
Loads a block into the buffer pool.
If no available slots are available, a block is evicted by LRU.
Offset is the byte offset of the block to be loaded.
*/
func (b *BufferPool) load(offset int64) *BlockElement {
	// Create new block
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
		b.flush(blockElement.Block)
	}

	// Add block to LRU queue
	block := NewBlock(bytes, offset)
	blockElement := &BlockElement{Block: block}
	blockElement.listElement = b.List.PushFront(blockElement)
	b.Blocks[block.Offset] = blockElement
	return blockElement
}

func (b *BufferPool) IsEmpty() bool {
	return b.List.Len() == 0
}

func (b *BufferPool) flushAll() {
	for item := b.List.Front(); b.List.Len() != 0; item = b.List.Front() {
		b.flush(item.Value.(*BlockElement).Block)
	}
}

func (b *BufferPool) flush(block *Block) {
	b.File.WriteAt(block.Bytes, block.Offset)
}

func (b *BufferPool) Shutdown() {
	b.flushAll()
	b.File.Close()
}
