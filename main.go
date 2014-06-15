package main

import (
	"fmt"
)

func main() {
	path := "./Test100 Copy.dat"

	bufferPool := NewBufferPool(path, 4, 4096, 3)
	defer bufferPool.Shutdown()
	rtn := bufferPool.GetRecord(0)
	fmt.Println(rtn)
	//bufferPool.Write(block)
	//bufferPool.Get(&rtn, 4)
	rtn = bufferPool.GetRecord(4096)
	fmt.Println(rtn)
	//bufferPool.Write(&rtn, 0)
	rtn = bufferPool.GetRecord(0)
	fmt.Println(rtn)
}
