/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/27 1:38 下午
 */
package test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

//探究 go 1.16 slice 切片扩容策略

//[0 ->    0] cap = 0     |  after append 1     cap = 1
//[0 ->    1] cap = 1     |  after append 2     cap = 2
//[0 ->    2] cap = 2     |  after append 3     cap = 4
//[0 ->    4] cap = 4     |  after append 5     cap = 8
//[0 ->    8] cap = 8     |  after append 9     cap = 16
//[0 ->   16] cap = 16    |  after append 17    cap = 32
//[0 ->   32] cap = 32    |  after append 33    cap = 64
//[0 ->   64] cap = 64    |  after append 65    cap = 128
//[0 ->  128] cap = 128   |  after append 129   cap = 256
//[0 ->  256] cap = 256   |  after append 257   cap = 512
//[0 ->  512] cap = 512   |  after append 513   cap = 1024
//[0 -> 1024] cap = 1024  |  after append 1025  cap = 1280
//[0 -> 1280] cap = 1280  |  after append 1281  cap = 1696
//[0 -> 1696] cap = 1696  |  after append 1697  cap = 2304
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

//append 的扩容 当添加的是一个切片时
func TestAppend(t *testing.T) {

	te := make([]int, 0)
	i := []int{1, 2}
	fmt.Println(te)
	i = append(i, 3, 4, 5)
	fmt.Println(len(i), cap(i))

}
func TestA(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	fmt.Println(s1)
	fmt.Println(s2)

}

func TestRange(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, v[i])
		fmt.Println(v[i])
	}
}

func TestAC(t *testing.T) {
	a := [2]int{0, 0}
	ad(a)
	fmt.Println(a)
}

func ad(a [2]int) {
	a[0] = 1
}

func TestBuilder(t *testing.T) {
	var list strings.Builder
	list.WriteString("11111\n")
	list.WriteString("22222\n")
	fmt.Println(list.String())
}

type info struct {
	S    string
	Name string
	P    string
	Q    string
	T    string
	LL   time.Time
}

func TestTrunCate(t *testing.T) {
	var s []info
	for i := 1; i < 6000000; i++ {
		s = append(s, info{
			S:    "123",
			Name: "231",
			P:    "3213",
			Q:    "3213",
			T:    "31123",
			LL:   time.Now(),
		})
	}

	fmt.Println(cap(s), len(s))

	s = s[0:0]
	s = nil
	fmt.Println(cap(s), len(s))
}

const s = "Go101.org"

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func TestAB(t *testing.T) {
	fmt.Println(1<<len(s[:])/128, b)
}

func TestDataTrace(t *testing.T) {
	var (
		a int         = 0
		b int64       = 0
		c interface{} = int(0)
		d interface{} = int64(0)
	)

	println(c == 0)
	println(c == a)
	println(c == b)
	println(d == b)
	println(d == 0)
}
