package s707

import "fmt"

type SingleNode struct {
	Val  int         // 节点的值
	Next *SingleNode // 下一个节点的指针
}

type MyLinkedList struct {
	dummyHead *SingleNode // 虚拟头节点
	Size      int         // 链表大小
}

func Exec() {
	list := Constructor()     // 初始化链表
	list.AddAtHead(100)       // 在头部添加元素
	list.AddAtTail(242)       // 在尾部添加元素
	list.AddAtTail(777)       // 在尾部添加元素
	list.AddAtIndex(1, 99999) // 在指定位置添加元素
	list.PrintLinkedList()    // 打印链表
}

// Constructor
// @Description 构造链表
// @Author lqc 2024-01-11 17:28:04
// @Return MyLinkedList
func Constructor() MyLinkedList {
	newNode := &SingleNode{ // 创建新节点
		-999,
		nil,
	}
	return MyLinkedList{ // 返回链表
		dummyHead: newNode,
		Size:      0,
	}
}

// Get
// @Description 获取节点数据
// @Author lqc 2024-01-11 17:28:13
// @Param index
// @Return int
func (list *MyLinkedList) Get(index int) int {
	/*if this != nil || index < 0 || index > this.Size {
		return -1
	}*/
	if list == nil || index < 0 || index >= list.Size { // 如果索引无效则返回-1
		return -1
	}
	// 让cur等于真正头节点
	cur := list.dummyHead.Next   // 设置当前节点为真实头节点
	for i := 0; i < index; i++ { // 遍历到索引所在的节点
		cur = cur.Next
	}
	return cur.Val // 返回节点值
}

// AddAtHead
// @Description 头插法
// @Author lqc 2024-01-11 17:28:23
// @Param val
func (list *MyLinkedList) AddAtHead(val int) {
	// 以下两行代码可用一行代替
	// newNode := new(SingleNode)
	// newNode.Val = val
	newNode := &SingleNode{Val: val}   // 创建新节点
	newNode.Next = list.dummyHead.Next // 新节点指向当前头节点
	list.dummyHead.Next = newNode      // 新节点变为头节点
	list.Size++                        // 链表大小增加1
}

// AddAtTail
// @Description  尾插法
// @Author lqc 2024-01-11 17:28:33
// @Param val
func (list *MyLinkedList) AddAtTail(val int) {
	newNode := &SingleNode{Val: val} // 创建新节点
	cur := list.dummyHead            // 设置当前节点为虚拟头节点
	for cur.Next != nil {            // 遍历到最后一个节点
		cur = cur.Next
	}
	cur.Next = newNode // 在尾部添加新节点
	list.Size++        // 链表大小增加1
}

// AddAtIndex
// @Description 插入指定位置
// @Author lqc 2024-01-11 17:28:49
// @Param index
// @Param val
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 { // 如果索引小于0，设置为0
		index = 0
	} else if index > list.Size { // 如果索引大于链表长度，直接返回
		return
	}

	newNode := &SingleNode{Val: val} // 创建新节点
	cur := list.dummyHead            // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ {     // 遍历到指定索引的前一个节点
		cur = cur.Next
	}
	newNode.Next = cur.Next // 新节点指向原索引节点
	cur.Next = newNode      // 原索引的前一个节点指向新节点
	list.Size++             // 链表大小增加1
}

// DeleteAtIndex
// @Description 删除指定位置
// @Author lqc 2024-01-11 17:29:09
// @Param index
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.Size { // 如果索引无效则直接返回
		return
	}
	cur := list.dummyHead        // 设置当前节点为虚拟头节点
	for i := 0; i < index; i++ { // 遍历到要删除节点的前一个节点
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next // 当前节点直接指向下下个节点，即删除了下一个节点
	}
	list.Size-- // 注意删除节点后应将链表大小减一
}

// PrintLinkedList
// @Description 打印链表元素
// @Author lqc 2024-01-11 17:29:25
func (list *MyLinkedList) PrintLinkedList() {
	cur := list.dummyHead // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}
