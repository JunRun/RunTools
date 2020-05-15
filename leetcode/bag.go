/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/13 2:09 下午
 */
package leetcode

import (
	"fmt"
)

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

// 动态规划 硬币问题
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1
func Knapsack3(coins []int, amount int) int {
	if amount < 0 || len(coins) < 1 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	minCoins := coins[0]
	//找出最小值
	for i := 1; i < len(coins); i++ {
		if minCoins > coins[i] {
			minCoins = coins[i]
		}
	}
	//计算层数
	level := amount / minCoins
	if level == 0 {
		return -1
	}
	//二维数组初始化
	state := make([][]bool, level)
	for i := range state {
		state[i] = make([]bool, amount+1)
	}
	for k := 0; k < len(coins); k++ {
		if coins[k] == amount {
			return 1
		}
		if coins[k] < amount {
			state[0][coins[k]] = true
		}
	}
	//状态赋值
	state[0][0] = true
	min := -1
	for i := 1; i < level; i++ {
		for j := 0; j <= amount; j++ {
			for k := 0; k < len(coins); k++ {
				if state[i-1][j] == true {
					if j+coins[k] == amount {
						min = i + 1
						return min
					}
					if j+coins[k] < amount {
						state[i][j+coins[k]] = true
					}
				}
			}
		}
	}
	//寻找最少的硬币组合个数
	for i := 0; i < level; i++ {
		if state[i][amount] == true {
			min = i + 1
			break
		}
	}
	return min
}

// 解 二
func Knapsack4(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}
			//+1 统计最小次数
			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}
