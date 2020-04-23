/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/23 4:23 下午
 */
package cache

import "testing"

type String string

func (s String) Len() int {
	return len(s)
}
func TestLru(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("k1", String("123"))
	if v, ok := lru.Get("k1"); !ok || string(v.(String)) != "123" {
		t.Fatalf("cache = 123 failed")
	}

}
