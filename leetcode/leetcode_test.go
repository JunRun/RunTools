/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/13 2:42 下午
 */
package leetcode

import "testing"

func TestBag(t *testing.T) {
	p := []int{10, 2, 4, 6, 10, 10, 27}
	Knapsack(p, len(p), 50)
}
