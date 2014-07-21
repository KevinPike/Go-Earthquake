package main

import (
	"container/list"
	"os"
)

type BufferPool struct {
	File       *os.File
	Blocks     map[int64]*BlockElement
	List       *list.List
	RecordSize int64
	BlockSize  int64
	Capacity   int
	Report     *Report
}

type BlockElement struct {
	Block       *Block
	listElement *list.Element
}

func NewBufferPool(path string, recordSize, blockSize, size int, report *Report) *BufferPool {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
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
		Report:     report,
	}
}

func (b *BufferPool) GetRecord(index int64) *Record {
	block, indexInBlock := b.getBlock(index)

	rtn := make([]byte, b.RecordSize)
	block.Get(&rtn, indexInBlock)

	record := NewRecord(rtn[0:2], rtn[2:4])
	return record
}

func (b *BufferPool) WriteRecord(record *Record, index int64) {
	block, indexInBlock := b.getBlock(index)

	bytes := record.ToBytes()

	block.Put(bytes, indexInBlock)
}

func (b *BufferPool) getBlock(recordIndex int64) (*Block, int64) {
	var totalOffsetInFile int64 = recordIndex * b.RecordSize

	blockOffset := totalOffsetInFile / b.BlockSize * b.BlockSize
	indexInBlock := totalOffsetInFile % b.BlockSize

	// Check if block exists in cache
	blockElement, exists := b.Blocks[blockOffset]
	if exists == false {
		blockElement = b.load(blockOffset)
		b.Report.CacheMisses++
	} else {
		b.Report.CacheHits++
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
	b.Report.DiskReads++

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
	for item := b.List.Front(); item != nil; item = item.Next() {
		b.flush(item.Value.(*BlockElement).Block)
	}
}

func (b *BufferPool) flush(block *Block) {
	b.File.WriteAt(block.Bytes, block.Offset)
	b.Report.DiskWrites++
}

func (b *BufferPool) Shutdown() {
	b.flushAll()
	b.File.Close()
}
