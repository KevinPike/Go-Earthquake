package main

func main() {
	path := "./Test100 Copy.dat"

	report := NewReport(path)

	bufferHeap := NewBufferHeap(path, 4, 4096, 20, report)
	defer bufferHeap.Shutdown()

	bufferHeap.Sort()

	report.Print()
}
