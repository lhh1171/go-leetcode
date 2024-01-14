package s236

import "tree/base"

// LowestCommonAncestor
// @Description		二叉树的最近公共祖先
// @Author lqc 2024-01-14 17:28:35
// @Param root
// @Param p
// @Param q
// @Return *base.TreeNode
func LowestCommonAncestor(root, p, q *base.TreeNode) *base.TreeNode {
	// check
	if root == nil {
		return nil
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
