// Package s209
// @Description:
package s209

// MinSubArrayLen
// @Description 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// @Author lqc 2024-01-11 14:41:41
// @Param target 目标和
// @Param nums 原数组
// @Return int 数组长度
func MinSubArrayLen(target int, nums []int) int {
	//  1.定义双指针
	left, length := 0, len(nums)
	// 子数组之和
	sum := 0
	// 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	result := length + 1
	for j := 0; j < length; j++ {
		sum += nums[j]
		// 刚刚大于目标和的时候进行循环，找出最小长度
		for sum >= target {
			subLength := j - left + 1

			// 取最小的长度
			if subLength < result {
				result = subLength
			}
			sum -= nums[left]
			left++
		}
	}
	if result == length+1 {
		return 0
	} else {
		return result
	}
}
