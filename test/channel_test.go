/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/28 5:16 下午
 */
package test

import (
	"fmt"
	"sync"
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

//串行化操作 并发访问slice
type serviceData struct {
	ch   chan int
	Data []int
}

func (s *serviceData) Schedule() {

	for i := range s.ch {
		s.Data = append(s.Data, i)
	}
}
func (s *serviceData) Close() {
	close(s.ch)
}
func (s *serviceData) addData(v int) {
	s.ch <- v
}

func NewScheduleJob(size int, done func()) *serviceData {
	s := &serviceData{
		ch:   make(chan int, size),
		Data: make([]int, 0),
	}
	go func() {
		s.Schedule()
		done()
	}()
	return s
}

//通过chan  来控制对slice 的并发写入操作
func TestCHM(t *testing.T) {
	var (
		wg sync.WaitGroup
		n  = 100
	)
	c := make(chan struct{})
	s := NewScheduleJob(n, func() { c <- struct{}{} })
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			s.addData(v)
		}(i)
	}
	wg.Wait()
	s.Close()
	<-c
	fmt.Println(len(s.Data))
}
func TestL(t *testing.T) {
	sl := make(chan int, 20)
	for i := 0; i < 40; i++ {
		sl <- i
		go func(i int) {
			fmt.Println(len(sl))
			fmt.Println(i)
			<-sl
		}(i)
	}

}

//协程池测试

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}
func (t *Task) Execute() {
	t.f()
}

type gPool struct {
	WorkNum int
	JobChan chan *Task //
	Entry   chan *Task //接收任务入口
}
type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func TestWO(t *testing.T) {

	var w worker = person{}
	fmt.Println(w)
}
