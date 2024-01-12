// Package s977
// @Description:
package s977

// SortedSquares
// @Description 有序数组（有正负数）,返回有序数组元素平方
// @Author lqc 2024-01-11 14:16:46
// @Param nums 原数组
// @Return []int 平方的数组
func SortedSquares(nums []int) []int {
	n := len(nums)
	//  1.取左边指针和右边指针，需要放置的位置（从后往前面放置）
	left, right, k := 0, n-1, n-1
	//  2.创建结果数组
	ans := make([]int, n)

	for left <= right {
		lm, rm := nums[left]*nums[left], nums[right]*nums[right]
		//  3.左边指针和右边指针代表元素的平方进行比较，谁比较大就插入在新数组的后面
		if lm > rm {
			ans[k] = lm
			left++
		} else {
			ans[k] = rm
			right--
		}
		k--
	}
	return ans
}
