/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-03 09:55
 */
package rdata

func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func MergeSort(a []int) []int {
	return mergeSortArray(a)
}

func mergeSortArray(a []int) []int {

	if len(a) <= 1 {
		return a
	}
	//取中间数
	m := len(a) / 2
	left := mergeSortArray(a[:m])
	right := mergeSortArray(a[m:])
	return merge(left, right)
}

func merge(left []int, right []int) (result []int) {
	l, r := 0, 0
	k, p := len(left), len(right)
	//注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < k && r < p {
		//从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
			//因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		} else {
			result = append(result, left[l])
			//因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）。
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
