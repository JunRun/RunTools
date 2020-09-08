package lock_free

import (
	"sync/atomic"
	"unsafe"
)

type LQueue struct {
	Head unsafe.Pointer
	Tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLQueue() *LQueue {
	n := unsafe.Pointer(&node{})
	return &LQueue{Head: n, Tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *LQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.Tail)
		next := load(&tail.next)
		if tail == load(&q.Tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&q.Tail, tail, n)
					return
				}
			} else {
				cas(&q.Tail, tail, n)
			}
		}
	}
}

//func (q *LQueue) Dequeue()interface{}{
//	for  {
//		head:=load(&q.Head)
//		tail:=load(&q.Tail)
//
//	}
//}
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
