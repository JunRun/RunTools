package main

import (
	"fmt"
	"sync"
)

// go trace 的使用 和go1.14 的异步抢占式调用
// go build -gcflags "-N -l" ./example.go
// go tool trace -http="127.0.0.1:6060" ./trace.output
func main() {
	SmallAllocation()
}

//go:noinline
func SmallAllocation() *smallStruct {
	return &smallStruct{}
}

type smallStruct struct {
	a, b int64
	c, d float64
}

func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}
