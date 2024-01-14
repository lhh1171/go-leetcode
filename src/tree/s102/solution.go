package s102

import (
	"container/list"
	"tree/base"
)

// LevelOrder1
// @Description  二叉树的递归遍历
// @Author lqc 2024-01-14 15:43:43
// @Param root
// @Return [][]int
func LevelOrder1(root *base.TreeNode) [][]int {
	res := [][]int{}

	depth := 0

	var order func(root *base.TreeNode, depth int)

	order = func(root *base.TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return res
}

// LevelOrder2
// @Description  二叉树的层序遍历  使用队列
// @Author lqc 2024-01-14 15:43:55
// @Param root
// @Return [][]int
func LevelOrder2(root *base.TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)

	var tmpArr []int

	for queue.Len() > 0 {
		//  保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*base.TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		//  放入结果集
		res = append(res, tmpArr)
		//  清空层的数据
		tmpArr = []int{}
	}

	return res
}

// LevelOrder3
// @Description  二叉树的层序遍历：使用切片模拟队列，易理解
// @Author lqc 2024-01-14 15:44:09
// @Param root
// @Return res
func LevelOrder3(root *base.TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	curLevel := []*base.TreeNode{root} // 存放当前层节点
	for len(curLevel) > 0 {
		nextLevel := []*base.TreeNode{} // 准备通过当前层生成下一层
		vals := []int{}

		for _, node := range curLevel {
			vals = append(vals, node.Val) // 收集当前层的值
			// 收集下一层的节点
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, vals)
		curLevel = nextLevel // 将下一层变成当前层
	}

	return
}
