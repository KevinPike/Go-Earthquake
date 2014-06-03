package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/Users/kevinpike/Development/Go/src/github.com/kevinpike/earthquake/Test100.dat"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := make([]byte, 4)
	_, err = file.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
	bufferPool := NewBufferPool(path, 5)
	rtn := make([]byte, 4)
	bufferPool.Get(&rtn, 0)
	fmt.Println(rtn)
	bufferPool.Get(&rtn, 1024)
	fmt.Println(rtn)
}
