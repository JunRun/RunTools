/**
 *
 * @Description: sync.Pool 学习测试
 * @Version: 1.0.0
 * @Date: 2020/4/21 9:36 上午
 */
package syncpool_test

import (
	"fmt"
	"sync"
	"testing"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func InitPool() {
	i := 0
	pool = &sync.Pool{
		New: func() interface{} {
			i++
			fmt.Println("初始化person", i)
			return new(Person)
		},
	}
}

func TestPoo(t *testing.T) {
	InitPool()

	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取", p)

	p.Name = "first"

	fmt.Println("设置 p.Name = ", p.Name)
	p = nil
	pool.Put(p)
	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}
