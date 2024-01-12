package s203

// ListNode
// @Description:单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList
// @Description 初始化链表
// @Author lqc 2024-01-11 16:51:55
// @Param arr 数组
func CreateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head *ListNode
	var cur = head
	for i := range arr {
		newNode := &ListNode{}
		newNode.Val = arr[i]
		if cur == nil {
			head = newNode
			cur = head
		} else {
			cur.Next = newNode
			cur = cur.Next
		}
	}

	return head
}

// RemoveElements1
// @Description 添加虚节点方式,时间复杂度 O(n)，空间复杂度 O(1)
// @Author lqc 2024-01-11 16:35:18
// @Param head       链表头结点
// @Param val        需要删除的值
// @Return *ListNode 返回的链表
func RemoveElements1(head *ListNode, val int) *ListNode {
	//  1.链表为空，不进行操作
	if head == nil {
		return head
	}
	//  2.定义虚拟节点
	dummy := &ListNode{}
	pre := dummy
	cur := head

	//  3.遍历链表
	for cur != nil {
		if cur.Val == val {
			//  进行删除
			pre.Next = cur.Next
		} else {
			// 正常遍历
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}

// RemoveElements2
// @Description 不添加虚拟节点and pre Node方式
// @Author lqc 2024-01-11 16:43:38
// @Param  head       链表头结点
// @Param  val        需要删除的值
// @Return *ListNode  返回的链表
func RemoveElements2(head *ListNode, val int) *ListNode {
	//  1.当head就是需要删除的点，进行删除
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head
	//  2.遍历链表
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == val {
			// 进行删除
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}

	return head
}
