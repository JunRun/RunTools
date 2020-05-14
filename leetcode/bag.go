/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/13 2:09 下午
 */
package leetcode

import "fmt"

//背包问题 动态规划求最优解(一维数组)

func Knapsack(p []int, n int, w int) {
	state := make([]bool, w+1)
	state[0] = true
	for i := 1; i < n; i++ {
		for j := w - p[i]; j >= 0; j-- {
			if state[j] == true { //放入背包
				state[j+p[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if state[i] == true {
			fmt.Println(i)
		}
	}

}

func Knapsack3(p []int, n, value, w int) {

}
