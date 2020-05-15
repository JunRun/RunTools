/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/15 11:04 上午
 */
package leetcode

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//失败 不能解决 如果有某一子数组 超了k 值，就会返回0
func SubarraySum(nums []int, k int) int {
	if len(nums) < 0 || len(nums) > 20000 {
		return 0
	}
	count := 0
	amount := k

	for i, j := 0, 0; i < len(nums); i++ {

		amount -= nums[i]
		if amount == 0 {
			count++
			amount = 0
			//如果 i+1 位 不与第 j位相同 那么 j+1-i+1 也不是子数组
			if i+1 < len(nums) && nums[i+1] == nums[j] {
				j = i + 1
				count++
			} else {
				j = i
			}
		}

	}
	return count
}

//思路：本质是寻找 j-i (0<j<i) 数组的和 等于 k
// 可以表示为 p[j...i]=k,若使 k[i]=0到 i p数组的和 可知 p[j...i]=k[i]-k[j]
// 求出数组中有多少 k[i]-k[j] == k 可以转换为 k[j]=k[i]-k

func SubarraySum2(nums []int, k int) int {
	res, sum := 0, 0
	tmp := make(map[int]int, 0)
	tmp[0] = 1
	for _, v := range nums {
		sum += v
		s := sum - k //k[j]=k[i]-k
		if _, ok := tmp[s]; ok {
			res += tmp[s]
		}
		tmp[sum]++
	}
	return res
}
