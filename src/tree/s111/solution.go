package s111

import "tree/base"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinDepth1
// @Description  递归
// @Author lqc 2024-01-14 16:59:37
// @Param root
// @Return int
func MinDepth1(root *base.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + MinDepth1(root.Right)
	}
	if root.Right == nil && root.Left != nil {
		return 1 + MinDepth1(root.Left)
	}
	return min(MinDepth1(root.Left), MinDepth1(root.Right)) + 1
}

// MinDepth2
// @Description  迭代
// @Author lqc 2024-01-14 16:59:44
// @Param root
// @Return int
func MinDepth2(root *base.TreeNode) int {
	dep := 0
	queue := make([]*base.TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			//  当遇见叶子节点时候返回
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}
