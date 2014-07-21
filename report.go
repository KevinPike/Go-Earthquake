package main

import (
	"fmt"
	"strings"
	"time"
)

type Report struct {
	File        string
	CacheHits   int
	CacheMisses int
	DiskReads   int
	DiskWrites  int
	startTime   time.Time
	endTime     time.Time
}

func NewReport(file string) *Report {
	return &Report{file, 0, 0, 0, 0, time.Now(), time.Now()}
}

func (report *Report) Start() {
	report.startTime = time.Now()
}

func (report *Report) End() {
	report.endTime = time.Now()
}

func (report *Report) Print() {
	fileParts := strings.Split(report.File, "/")
	fileName := fileParts[len(fileParts)-1]
	fmt.Printf("Sorting \"%s\"\n", fileName)
	fmt.Println("Number of cache hits:", report.CacheHits)
	fmt.Println("Number of cache misses:", report.CacheMisses)
	fmt.Println("Number of disk reads:", report.DiskReads)
	fmt.Println("Number of disk writes:", report.DiskWrites)
	fmt.Println("Time to execute the heapsort:", report.endTime.Sub(report.startTime))
}
