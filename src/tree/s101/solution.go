package s101

import "tree/base"

// 递归
func defs(left *base.TreeNode, right *base.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}

// IsSymmetric1
// @Description 是否是对称二叉数 递归
// @Author lqc 2024-01-14 16:37:12
// @Param root
// @Return bool
func IsSymmetric1(root *base.TreeNode) bool {
	return defs(root.Left, root.Right)
}

// IsSymmetric2
// @Description	迭代
// @Author lqc 2024-01-14 16:37:32
// @Param root
// @Return bool
func IsSymmetric2(root *base.TreeNode) bool {
	var queue []*base.TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}
