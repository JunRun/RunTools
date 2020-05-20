/**
 *
 * Description: 数据结构 堆的学习
 * Version: 1.0.0
 * Date: 2020/5/18 3:07 下午
 */
package heap

import "fmt"

type heap struct {
	a     []int
	n     int
	count int
}

func Heap(capacity int) *heap {
	return &heap{
		a: make([]int, capacity),
		n: capacity,
	}
}

func (h *heap) Insert(data int) bool {
	if h.count > h.n {
		fmt.Println("堆满了")
		return false
	}
	h.count++

	h.a[h.count] = data
	for i := h.count; i/2 > 0 && h.a[i] > h.a[i/2]; i = i / 2 {
		swap(h.a, i, i/2)
	}
	return true
}

func swap(a []int, pre, po int) {
	t := a[pre]
	a[pre] = a[po]
	a[po] = t
}

//从上往下堆化

func (h *heap) buildHeap() {
	for i := h.count / 2; i > 0; i-- {
		heapify(h.a, h.count, i)
	}
}
func heapify(a []int, n, i int) {
	for {
		maxPoi := i
		if a[i] < a[i*2] && i*2 <= n {
			maxPoi = i * 2
		}
		if i*2+1 <= n && a[maxPoi] < a[i*2+1] {
			maxPoi = i*2 + 1
		}
		if maxPoi == i {
			break
		}
		swap(a, maxPoi, i)
		i = maxPoi
	}

}
