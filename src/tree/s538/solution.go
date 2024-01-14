package s538

import (
	"tree/base"
)

var pre int

// ConvertBST
// @Description
// @Author lqc 2024-01-14 17:45:08
// @Param root
// @Return *base.TreeNode
func ConvertBST(root *base.TreeNode) *base.TreeNode {
	pre = 0
	traversal(root)
	return root
}

func traversal(cur *base.TreeNode) {
	if cur == nil {
		return
	}
	//  递归右边节点
	traversal(cur.Right)
	cur.Val += pre
	pre = cur.Val
	//  节点的左杰点是该节点的值（新值）加上左杰点的值
	traversal(cur.Left)
}
