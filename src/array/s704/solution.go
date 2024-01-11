// Package s704
// @Description:
package s704

import "strconv"

// Search1 @Title Search1
// @Description （版本一）左闭右闭区间
// @Author lqc 2024-01-11 10:59:35
// @Param nums 被查询的数组
// @Param target 查询目标数
// @Return int 返回查询下标
func Search1(nums []int, target int) int {
	// 1.定义不变量
	high := len(nums) - 1
	low := 0
	// 2.当满足条件时，进行循环
	for low <= high {
		//  3.每次取mid的时候，偏左一点
		mid := low + (high-low)/2
		println("low取的是  " + strconv.Itoa(low))
		println("high取的是  " + strconv.Itoa(high))
		println("mid取的是  " + strconv.Itoa(mid))
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// Search2
// @Description （版本二）左闭右开区间
// @Author lqc 2024-01-11 11:01:05
// @Param nums 被查询的数组
// @Param target 查询目标数
// @Return int 返回查询下标
func Search2(nums []int, target int) int {
	high := len(nums)
	low := 0
	for low < high {
		mid := low + (high-low)/2
		println("low取的是  " + strconv.Itoa(low))
		println("high取的是  " + strconv.Itoa(high))
		println("mid取的是  " + strconv.Itoa(mid))
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
