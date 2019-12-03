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

func MergeSort(a []int, len int) {
	mergeSortArray(a, 0, len-1)
}

func mergeSortArray(a []int, l int, r int) {

	if l >= r {
		return
	}
	//取中间数
	m := (l + r) / 2
	mergeSortArray(a, l, m)
	mergeSortArray(a, m+1, r)
	merge(a, l, m, r)
}

func merge(a []int, l int, m int, r int) {
	temp := make([]int, r-l)

	i := l
	j := m + 1
	k := 0
	for i < m && j < r {
		if a[i] <= a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
			temp[k] = a[j]
			k++
			j++
		}
	}
	//判断有无剩余数据
	start := i
	end := m
	if j <= r {
		start = j
		end = r
	}
	for ; start <= end; start++ {
		temp[k] = a[start]
	}
	//将临时数组的值拷贝回原数组
	for x := 0; x < (r - l); x++ {
		a[l+x] = temp[x]
	}
}
