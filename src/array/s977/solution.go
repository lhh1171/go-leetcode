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
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > rm {
			ans[k] = lm
			i++
		} else {
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}
