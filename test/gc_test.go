/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/27 2:30 下午
 */
package test

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

func TestGc(t *testing.T) {
	go func() {
		for {

		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	fmt.Println("ok")
}

// go trace 的使用 和go1.14 的异步抢占式调用
func TestTrace(t *testing.T) {
	runtime.GOMAXPROCS(1)
	create, _ := os.Create("trace.output")
	defer create.Close()
	_ = trace.Start(create)
	defer trace.Stop()
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(&w, i)
	}
	wg.Wait()
}
func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}
