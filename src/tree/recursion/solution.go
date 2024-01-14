package recursion

import "tree/base"

// PreorderTraversal
// @Description 前序遍历
// @Author lqc 2024-01-14 11:26:16
// @Param root
// @Return res
func PreorderTraversal(root *base.TreeNode) (res []int) {
	//  构造函数指针
	var traversal func(node *base.TreeNode)

	//  初始化函数，匿名函数
	traversal = func(node *base.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// InorderTraversal
// @Description  中序遍历
// @Author lqc 2024-01-14 11:26:28
// @Param root
// @Return res
func InorderTraversal(root *base.TreeNode) (res []int) {
	var traversal func(node *base.TreeNode)
	traversal = func(node *base.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// PostorderTraversal
// @Description 后序遍历
// @Author lqc 2024-01-14 11:26:52
// @Param root
// @Return res
func PostorderTraversal(root *base.TreeNode) (res []int) {
	var traversal func(node *base.TreeNode)
	traversal = func(node *base.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
