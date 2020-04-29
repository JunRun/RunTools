/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/28 5:16 下午
 */
package test

import (
	"fmt"
	"testing"
)

//通过多线程的方式找出 1000以内的素数
func TestChannle(t *testing.T) {
	option, wait := make(chan int), make(chan struct{})
	ProcessOn(option, wait)
	for i := 2; i < 1000; i++ {
		option <- i
	}
	close(option)
	<-wait
}

func ProcessOn(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		ProcessOn(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
	}()
}
