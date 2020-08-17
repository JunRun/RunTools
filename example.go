package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

// go trace 的使用 和go1.14 的异步抢占式调用
// go build -gcflags "-N -l" ./example.go
// go tool trace -http="127.0.0.1:6060" ./trace.output
func main() {
	runtime.GOMAXPROCS(2)
	create, _ := os.Create("trace.output")
	defer create.Close()
	_ = trace.Start(create)
	defer trace.Stop()
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go calcSum(&w, i)
	}
	w.Wait()
}
func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}
