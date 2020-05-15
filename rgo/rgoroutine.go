/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/15 3:58 下午
 */
package rgo

import "sync"

var mu sync.Mutex
var chain string

func F() {
	chain = "main"
	A()
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "-->A"
	B()

}
func B() {
	chain = chain + "--->B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "---C"
}
