/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-26 13:33
 */
package leetcode

import "strings"

/**
leetCode 解题 找出没有重复字符的最长子串
*/
func LengthOfString(s string) int {
	var length, j, i, t = 0, 0, 0, 0
	for ; i < len(s); i++ {
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
