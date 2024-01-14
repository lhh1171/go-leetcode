package s104

import "tree/base"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxDepth1
// @Description  递归
// @Author lqc 2024-01-14 16:47:56
// @Param root
// @Return int
func MaxDepth1(root *base.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(MaxDepth1(root.Left), MaxDepth1(root.Right)) + 1
}

// MaxDepth2
// @Description  迭代遍历
// @Author lqc 2024-01-14 16:47:50
// @Param root
// @Return int
func MaxDepth2(root *base.TreeNode) int {
	levl := 0
	queue := make([]*base.TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			//  左边入栈
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			//  右边入栈
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			//  将第一个元素移除，此时的queue只是这一层的元素
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}
