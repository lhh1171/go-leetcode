package s108

import "tree/base"

// SortedArrayToBST
// @Description  递归
// @Author lqc 2024-01-14 17:41:54
// @Param nums
// @Return *base.TreeNode
func SortedArrayToBST(nums []int) *base.TreeNode {
	if len(nums) == 0 { //终止条件，最后数组为空则可以返回
		return nil
	}
	//  从中间开始构建
	idx := len(nums) / 2
	root := &base.TreeNode{Val: nums[idx]}

	root.Left = SortedArrayToBST(nums[:idx])
	root.Right = SortedArrayToBST(nums[idx+1:])

	return root
}
