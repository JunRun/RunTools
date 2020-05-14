/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/27 2:30 下午
 */
package test

import (
	"fmt"
	"runtime"
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
