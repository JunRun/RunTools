/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/23 9:56 上午
 */
package cache

import "container/list"

type RCache struct {
	maxBytes int64 //最大使用内存
	nbBytes  int64 //当前使用的内存
	ll       *list.List
	cache    map[string]*list.Element
	onE      func(key string, value Value)
}

type Entry struct {
	key   string
	Value Value
}
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *RCache {
	return &RCache{
		maxBytes: maxBytes,
		nbBytes:  0,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		onE:      onEvicted,
	}
}

func (r *RCache) Get(key string) (value Value, ok bool) {
	if ele, ok := r.cache[key]; ok {
		r.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		return kv.Value, true
	}
	return
}

func (r *RCache) Add(key string, value Value) {

	if ele, ok := r.cache[key]; ok { //若已存在键值，则更新
		r.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		r.nbBytes += int64(value.Len()) - int64(kv.Value.Len())
		kv.Value = value

	} else { //若不存在,则将存入

		ele := r.ll.PushFront(&Entry{
			key:   key,
			Value: value,
		})
		r.cache[key] = ele
		r.nbBytes += int64(len(key)) + int64(value.Len())
	}
	//循环判断 是否超出最大内存值，是则调用删除节点
	for r.maxBytes != 0 && r.maxBytes < r.nbBytes {
		r.RemoveOldest()
	}
}
func (r *RCache) RemoveOldest() {
	//取到首节点
	ele := r.ll.Back()
	if ele != nil {
		kv := ele.Value.(*Entry)
		//删除该节点的映射关系
		r.ll.Remove(ele)
		delete(r.cache, kv.key)
		r.nbBytes -= int64(len(kv.key)) - int64(kv.Value.Len())
		//如果回调函数不为空，则执行
		if r.onE != nil {
			r.onE(kv.key, kv.Value)
		}
	}
}

func (r *RCache) Len() int {
	return r.ll.Len()
}
