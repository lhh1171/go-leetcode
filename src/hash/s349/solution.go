package s349

// Intersection
// @Description 两个数组的交集
// @Author lqc 2024-01-11 19:39:43
// @Param nums1
// @Param nums2
// @Return []int
func Intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
