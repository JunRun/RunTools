/**
 *
 * @Description: 快速排序
 * @Version: 1.0.0
 * @Date: 2019-12-04 11:20
 */
package rdata

func QuickSort(a []int) {
	quickSortArray(a, 0, len(a)-1)
}

//递归函数
func quickSortArray(a []int, left int, right int) {
	if left > right {
		return
	}
	p := partition(a, left, right)
	quickSortArray(a, left, p-1)
	quickSortArray(a, p+1, right)

}

//
func partition(a []int, left int, right int) int {
	pivot := a[right]
	i := left
	for j := left; j <= right-1; j++ {
		if a[j] < pivot {
			t := a[i]
			a[i] = a[j]
			a[j] = t
			i++
		}
	}
	t := a[right]
	a[right] = a[i]
	a[i] = t
	return i
}
