/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-26 13:33
 */
package leetcode

import (
	"strings"
)

/**
leetCode 解题 找出没有重复字符的最长子串
*/
func LengthOfString(s string) int {
	var length, j, i, t = 0, 0, 0, 0
	for ; i < len(s); i++ {
		//返回字符 在字符串中的位置
		if t = strings.Index(s[j:i], string(s[i])); t == -1 {
			if length < i-j {
				length = i - j + 1
			}
		} else {
			j = j + t - 1
		}

	}
	return length
}

/**给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。
*/
func SearchRange(nums []int, target int) []int {
	var i []int
	start, end := -1, -1
	left, right := 0, len(nums)-1
	middle := 0

	for l := 0; l < len(nums); l++ {

		if left > right {
			break
		}
		//middle = int(math.Floor(float64((left + right) / 2)))
		middle = (left + right) / 2
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			break
		}
	}
	if left > right {
		return append(i, -1, -1)
	}
	start = middle
	for {
		start--
		if start < 0 {
			start++
			break
		}
		if nums[start] != target {
			start++
			break
		}

	}
	end = middle
	for {
		end++
		if end > len(nums)-1 {
			end--
			break
		}
		if nums[end] != target {
			end--
			break
		}
	}
	i = append(i, start, end)
	return i
}
