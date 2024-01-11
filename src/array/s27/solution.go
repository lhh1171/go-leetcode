// Package s27
// @Description:
package s27

// RemoveElement1
// @Description 暴力双重for循环
// @Author lqc 2024-01-11 14:05:09
// @Param nums 被删除的数组
// @Param val 删除的元素
// @Return int 返回新的长度
func RemoveElement1(nums []int, val int) int {
	length := len(nums)
	res := 0
	//  遍历数组，等于目标值，后面所有的元素都往前移动
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

// RemoveElement2 @Title RemoveElement2
// @Description  相向双指针法
// @Author lqc 2024-01-11 14:00:22
// @Param nums 被删除的数组
// @Param val 删除的元素
// @Return int 返回新的长度
func RemoveElement2(nums []int, val int) int {
	// 有点像二分查找的左闭右闭区间 所以下面是<=
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉

		// 不断将左边指针右移，直到等于目标值
		for left <= right && nums[left] != val {
			left++
		}

		//  右边指针遇见等于目标值就往前移动，知道不等于目标值
		for left <= right && nums[right] == val {
			right--
		}

		// 各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	//fmt.Println(nums)
	return left
}
