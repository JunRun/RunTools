/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/13 2:42 下午
 */
package leetcode

import (
	"fmt"
	"testing"
)

func TestBag(t *testing.T) {
	p := []int{10, 2, 4, 6, 10, 10, 27}
	Knapsack(p, len(p), 50)
}

func TestP(t *testing.T) {

	//[1,2147483647]
	//

	//[120,6,320,300,100,192,212,89,106,461]
	//8332
	//conis:=[]int{120,6,320,300,100,192,212,89,106,461}
	//fmt.Println(Knapsack3(conis,8332))

	//coins2:=[]int{1}
	//fmt.Println(Knapsack3(coins2,1))

	coins3 := []int{1, 5, 3}
	fmt.Println(Knapsack4(coins3, 11))
}
func TestA(t *testing.T) {
	fmt.Println(3 / 2)
}

func TestSubarraySum(t *testing.T) {
	//i:=[]int{-1,-1,1}
	//k:=0

	i := []int{1, 1, 4, 1, 1, 4}
	k := 2
	fmt.Println(SubarraySum2(i, k))
}

func TestConstructor(t *testing.T) {
	//a :=[]int{2,0,-1,1,-2,-4,3}
	s := []int{4, 5, 8, 2}
	larg := Constructor(3, s)
	fmt.Println(larg.Add(3))
	fmt.Println(larg.Add(1))
	fmt.Println(larg.Add(-2))
	fmt.Println(larg.Add(-4))
	fmt.Println(larg.Add(3))

}
