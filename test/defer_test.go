/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/20 3:19 下午
 */
package test

import (
	"fmt"
	"testing"
)

func TestDe(t *testing.T) {
	s := ""
	defer func(s string) {
		if s == "" {
			fmt.Println("s 为空")
		} else {
			fmt.Println("不为空")
		}
	}(s)
	s = "23"
}

func TestDeL(t *testing.T) {
	fmt.Println(deferPR())
	fmt.Println(foo())
}
func deferPR() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func foo() int {
	i := 1
	defer func() {
		i++
	}()

	return i
}
