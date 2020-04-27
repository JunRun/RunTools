/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/27 1:38 下午
 */
package test

import (
	"fmt"
	"testing"
)

//探究 go 1.4 slice 切片扩容策略
func TestSlice(t *testing.T) {

	s := make([]int, 0)
	oldCap := cap(s)
	for i := 1; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

func TestA(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	fmt.Println(s1)
	fmt.Println(s2)

}
