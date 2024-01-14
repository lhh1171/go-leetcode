package s106

import "tree/base"

var (
	hash map[int]int
)

// BuildTree
// 第一步：如果数组大小为零的话，说明是空节点了。
// 第二步：如果不为空，那么取后序数组最后一个元素作为节点元素。
// 第三步：找到后序数组最后一个元素在中序数组的位置，作为切割点
// 第四步：切割中序数组，切成中序左数组和中序右数组 （顺序别搞反了，一定是先切中序数组）
// 第五步：切割后序数组，切成后序左数组和后序右数组
// 第六步：递归处理左区间和右区间
func BuildTree(inorder []int, postorder []int) *base.TreeNode {
	hash = make(map[int]int)
	for i, v := range inorder { // 用map保存中序序列的数值对应位置
		hash[v] = i
	}
	// 以左闭右闭的原则进行切分
	root := rebuild(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
	return root
}

// rootIdx表示根节点在后序数组中的索引，l, r 表示在中序数组中的前后切分点
// rebuild
// @Description
// @Author lqc 2024-01-14 17:14:30
// @Param inorder
// @Param postorder
// @Param rootIdx 后续遍历的最后一个是根节点
// @Param index	截取中序遍历的位置
// @Param r
// @Return *base.TreeNode
func rebuild(inorder []int, postorder []int, rootIdx int, index, r int) *base.TreeNode {
	if index > r { // 说明没有元素，返回空树
		return nil
	}
	if index == r { // 只剩唯一一个元素，直接返回
		return &base.TreeNode{Val: inorder[index]}
	}
	rootV := postorder[rootIdx]        // 根据后序数组找到根节点的值
	rootIn := hash[rootV]              // 找到根节点在对应的中序数组中的位置
	root := &base.TreeNode{Val: rootV} // 构造根节点
	// 重建左节点和右节点
	root.Left = rebuild(inorder, postorder, rootIdx-(r-rootIn)-1, index, rootIn-1)
	root.Right = rebuild(inorder, postorder, rootIdx-1, rootIn+1, r)
	return root
}
