package main

import (
	"fmt"
)

func main() {
	path := "./Test100.dat"

	bufferPool := NewBufferPool(path, 4, 4096, 1)
	defer bufferPool.Shutdown()
	rtn := make([]byte, 4)
	bufferPool.Get(&rtn, 0)
	fmt.Println(rtn)
	bufferPool.Get(&rtn, 2)
	fmt.Println(rtn)
	bufferPool.Get(&rtn, 1024)
	fmt.Println(rtn)
	bufferPool.Get(&rtn, 1023)
}
