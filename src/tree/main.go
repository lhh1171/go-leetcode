package main

import (
	"tree/base"
	"tree/s538"
)

func main() {
	//  前序遍历：中左右
	//	中序遍历：左中右
	//	后序遍历：左右中
	//  递归遍历
	//fmt.Println(recursion.PreorderTraversal(&base.TreeNode{
	//	Val: 3,
	//}))

	//  二叉树的迭代遍历
	//fmt.Println(iteration.PreorderTraversal(&base.TreeNode{
	//	Val: 3,
	//}))

	//  102.二叉树的层序遍历	https://leetcode.cn/problems/binary-tree-level-order-traversal/
	//s102.LevelOrder1(&base.TreeNode{})
	//s102.LevelOrder2(&base.TreeNode{})

	//  翻转二叉树  https://leetcode.cn/problems/invert-binary-tree/ 遍历交换
	//s226.InvertTree1(&base.TreeNode{})

	//  对称二叉树	https://leetcode.cn/problems/symmetric-tree/
	//s101.IsSymmetric1(new(base.TreeNode))

	//  二叉树的最大深度	https://leetcode.cn/problems/maximum-depth-of-binary-tree/
	//s104.MaxDepth1(new(base.TreeNode))

	//  二叉树的最小深度	https://leetcode.cn/problems/minimum-depth-of-binary-tree/
	//s111.MinDepth1(new(base.TreeNode))

	//  二叉树的所有路径	https://leetcode.cn/problems/binary-tree-paths/
	//s257.BinaryTreePaths1(new(base.TreeNode))

	//  从中序与后序遍历序列构造二叉树	https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
	//s106.BuildTree([]int{}, []int{})

	//  二叉树的最近公共祖先	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
	//s236.LowestCommonAncestor(new(base.TreeNode), new(base.TreeNode), new(base.TreeNode))

	//  二叉搜索树的最近公共祖先	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
	//s235.LowestCommonAncestor(new(base.TreeNode), new(base.TreeNode), new(base.TreeNode))

	//  将有序数组转换为二叉搜索树	https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
	//s108.SortedArrayToBST([]int{})

	//  把二叉搜索树转换为累加树	https://leetcode.cn/problems/convert-bst-to-greater-tree/
	s538.ConvertBST(new(base.TreeNode))
}
