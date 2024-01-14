package s226

import (
	"container/list"
	"tree/base"
)

// InvertTree1
// @Description  前序遍历
// @Author lqc 2024-01-14 16:17:58
// @Param root
// @Return *base.TreeNode
func InvertTree1(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left //交换

	InvertTree1(root.Left)
	InvertTree1(root.Right)

	return root
}

// InvertTree2
// @Description  递归版本的后序遍历
// @Author lqc 2024-01-14 16:18:05
// @Param root
// @Return *base.TreeNode
func InvertTree2(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return root
	}

	InvertTree2(root.Left)                        //遍历左节点
	InvertTree2(root.Right)                       //遍历右节点
	root.Left, root.Right = root.Right, root.Left //交换

	return root
}

// InvertTree3
// @Description  迭代版本的前序遍历
// @Author lqc 2024-01-14 16:18:13
// @Param root
// @Return *base.TreeNode
func InvertTree3(root *base.TreeNode) *base.TreeNode {
	stack := []*base.TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left //交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return root
}

// InvertTree4
// @Description  迭代版本的后序遍历
// @Author lqc 2024-01-14 16:18:20
// @Param root
// @Return *base.TreeNode
func InvertTree4(root *base.TreeNode) *base.TreeNode {
	stack := []*base.TreeNode{}
	node := root
	var prev *base.TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			node.Left, node.Right = node.Right, node.Left //交换
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}

	return root
}

// InvertTree5
// @Description  层序遍历
// @Author lqc 2024-01-14 16:18:26
// @Param root
// @Return *base.TreeNode
func InvertTree5(root *base.TreeNode) *base.TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*base.TreeNode)
			e.Left, e.Right = e.Right, e.Left //交换
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}
