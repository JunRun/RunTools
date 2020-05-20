/**
 *
 * Description:
 * Version: 1.0.0
 * Date: 2020/5/19 11:06 上午
 */
package leetcode

type KthLargest struct {
	k     int
	nums  []int
	count int
}

// 创建 一个 长度  k 的小顶堆 ，构造时候进行堆化，堆顶元素 是
// 第k 大的 数字,添加元素后 ，大于对顶元素，移除对顶元素，进行堆化，然后返回堆顶元素。
func Constructor(k int, nums []int) KthLargest {
	large := KthLargest{
		k:    k,
		nums: make([]int, k+1),
	}
	for i := 0; i < len(nums); i++ {
		large.Add(nums[i])
	}
	return large
}

func (this *KthLargest) Add(val int) int {
	if this.count >= this.k {
		if this.nums[1]-val >= 0 {
			return this.nums[1]
		}
		this.nums[1] = val
	} else {
		this.nums[this.k] = val
	}
	this.count++

	for i := this.k / 2; i > 0; i-- {
		for {
			maxPoi := i
			if this.nums[i] > this.nums[i*2] && 2*i <= this.k {
				maxPoi = 2 * i
			}
			if 2*i+1 <= this.k && this.nums[maxPoi] > this.nums[i*2+1] {
				maxPoi = 2*i + 1
			}
			if maxPoi == i {
				break
			}
			temp := this.nums[i]
			this.nums[i] = this.nums[maxPoi]
			this.nums[maxPoi] = temp
		}

	}
	return this.nums[1]
}
