package s19

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

// RemoveNthFromEnd
// @Description 删除链表的倒数第N个节点  双指针法
// @Author lqc 2024-01-11 18:50:16
// @Param head
// @Param n 位置在倒数第n个
// @Return *ListNode
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
