package s257

import (
	"strconv"
	"tree/base"
)

func BinaryTreePaths1(root *base.TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *base.TreeNode, s string)
	//  定义匿名方法
	travel = func(node *base.TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			//  当左右孩子为空，添加到res
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return res
}

func BinaryTreePaths2(root *base.TreeNode) []string {
	stack := []*base.TreeNode{}
	paths := make([]string, 0)
	res := make([]string, 0)
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		path := paths[l-1]
		stack = stack[:l-1]
		paths = paths[:l-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			//  叠加路径
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			//  叠加路径
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}
