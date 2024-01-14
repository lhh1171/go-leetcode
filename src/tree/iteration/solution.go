package iteration

import (
	"container/list"
	"tree/base"
)

// PreorderTraversal
// @Description 迭代法前序遍历 利用栈
// @Author lqc 2024-01-14 11:49:43
// @Param root
// @Return []int
func PreorderTraversal(root *base.TreeNode) []int {
	var ans []int

	if root == nil {
		return ans
	}

	stack := list.New()
	//  中入栈
	stack.PushBack(root)

	for stack.Len() > 0 {
		//  出栈（先出战的是中，后面就是左右）
		node := stack.Remove(stack.Back()).(*base.TreeNode)

		//  计算结果
		ans = append(ans, node.Val)

		//  右入栈
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		//  左入栈
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ans
}

// PostorderTraversal
// @Description  先序遍历是中左右，后续遍历是左右中，那么我们只需要调整一下先序遍历的代码顺序，
//
//	就变成中右左的遍历顺序，然后在反转result数组，输出的结果顺序就是左右中了
//
// @Author lqc 2024-01-14 12:00:28
// @Param root
// @Return []int
func PostorderTraversal(root *base.TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*base.TreeNode)

		ans = append(ans, node.Val)

		//  左入栈
		if node.Left != nil {
			st.PushBack(node.Left)
		}

		//  右入栈
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}

	//  中右左
	//  左右中
	reverse(ans)
	return ans
}

// reverse
// @Description 反转
// @Author lqc 2024-01-14 11:57:42
// @Param a
func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}

// InorderTraversal
// @Description 中序遍历
// @Author lqc 2024-01-14 15:34:38
// @Param root
// @Return []int
func InorderTraversal(root *base.TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	st := list.New()
	cur := root

	//  栈为空且cur为空，停止循环
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			//  中入栈
			st.PushBack(cur)

			//  当前节点变成其左杰点
			cur = cur.Left
		} else {
			//  当当前节点为空时，弹出栈，空节点的父节点
			cur = st.Remove(st.Back()).(*base.TreeNode)
			//  拼装结果
			res = append(res, cur.Val)
			//  当前节点变成右节点
			cur = cur.Right
		}
	}

	return res
}
