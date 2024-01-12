package s454

// FourSumCount
// @Description 四数之和
// @Author lqc 2024-01-11 19:58:04
// @Param nums1
// @Param nums2
// @Param nums3
// @Param nums4
// @Return int
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int) //key:a+b的数值，value:a+b数值出现的次数
	count := 0
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}
	return count
}
