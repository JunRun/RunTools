/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-12 14:04
 */
package rgo

import (
	"fmt"
	"sync"
	"time"
)

var Group sync.WaitGroup

func Timer() {
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("time out...")
				Group.Done()
			}
		}
	}()
}
