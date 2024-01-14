package s235

import (
	"tree/base"
)

// LowestCommonAncestor
// @Description
// @Author lqc 2024-01-14 17:34:45
// @Param root
// @Param p
// @Param q
// @Return *base.TreeNode
func LowestCommonAncestor(root, p, q *base.TreeNode) *base.TreeNode {
	if root == nil {
		return nil
	}
	for {
		//  当前节点
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		//  直到该节点在p和q之间
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}

}
