package s160

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

// GetIntersectionNode1
// @Description 遍历法
// @Author lqc 2024-01-11 18:56:50
// @Param headA 链表a头节点
// @Param headB 链表b头节点
// @Return *ListNode 相交的第一个节点
func GetIntersectionNode1(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast.Val != slow.Val {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// GetIntersectionNode2
// @Description 双指针法
// @Author lqc 2024-01-11 18:56:16
// @Param headA 链表a头节点
// @Param headB 链表b头节点
// @Return *ListNode 相交的第一个节点
func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1.Val != l2.Val {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}

	return l1
}
