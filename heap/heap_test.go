/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/19 9:18 上午
 */
package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := Heap(30)

	a := []int{5, 2, 3, 4, 9, 6, 7, 30, 40, 10}
	fmt.Println(a)
	for _, v := range a {
		h.Insert(v)
	}
	fmt.Println(h.a)
	h.buildHeap()
	fmt.Println(h.a)
}
