package s1

// TwoSum1
// @Description  暴力解法
// @Author lqc 2024-01-11 19:53:20
// @Param nums
// @Param target
// @Return []int
func TwoSum1(nums []int, target int) []int {
	for k1, _ := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1]+nums[k2] {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}

// TwoSum2
// @Description  使用map方式解题，降低时间复杂度,把缺的哪一点放入map中
// @Author lqc 2024-01-11 19:53:13
// @Param nums
// @Param target
// @Return []int
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
