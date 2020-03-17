/**
 *
 * @Description: 连接池
 * @Version: 1.0.0
 * @Date: 2020-03-02 15:33
 */
package rsqlPool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	Lock    sync.Mutex
	Res     chan io.Closer
	Factory func() (io.Closer, error)
	Closed  bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size 值过小")
	}
	return &Pool{
		Res:     make(chan io.Closer, size),
		Factory: fn,
	}, nil
}

//获取连接
func (p *Pool) Get() (io.Closer, error) {
	select {
	case r, ok := <-p.Res:
		if !ok {
			return nil, errors.New("连接池关闭")
		}
		return r, nil
	default:
		fmt.Println("创建资源")
		return p.Factory()
	}
}

//关闭连接池
func (p *Pool) Close() {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	if p.Closed {
		return
	}
	p.Closed = true
	//关闭通道
	close(p.Res)
	//关闭资源
	for r := range p.Res {
		r.Close()
	}
}

//释放资源
func (p *Pool) Release(r io.Closer) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	if p.Closed {
		r.Close()
		return
	}

	select {
	case p.Res <- r:
		fmt.Println("释放资源，放入连接池")
	default:
		r.Close()
		fmt.Println("连接池已满，关闭资源")
	}
}
