package s24

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

// SwapPairs1
// @Description 虚节点版本
// @Author lqc 2024-01-11 17:51:48
// @Param head
// @Return *ListNode
func SwapPairs1(head *ListNode) *ListNode {
	//  1.创建虚节点
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy

	//  head不为空，head的Next不为空，就会进行交换
	for head != nil && head.Next != nil {
		//  虚指针指向第二个
		pre.Next = head.Next
		//  第二个指向头结点
		next := head.Next.Next
		head.Next.Next = head
		//  头节点指向第三个
		head.Next = next

		//  虚指针的next指向第三个
		pre = head
		//  头指针变成了第三个
		head = next
	}
	return dummy.Next
}

// SwapPairs2
// @Description 递归版本
// @Author lqc 2024-01-11 17:52:03
// @Param head
// @Return *ListNode
func SwapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = SwapPairs2(next.Next)
	next.Next = head
	return next
}
