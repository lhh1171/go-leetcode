package s142

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

// DetectCycle
// @Description 检测环链表
// @Author lqc 2024-01-11 19:12:07
// @Param head
// @Return *ListNode
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		//  慢指针每次走一步
		slow = slow.Next
		//  快指针每次走两步
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
