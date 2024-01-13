package s225

type MyStack struct {
	//创建两个队列
	queue1 []int
	queue2 []int
}

func Exec() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	println(stack.Pop())
	println(stack.Top())
}

func Constructor() MyStack {
	return MyStack{ //初始化
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	//先将数据存在queue2中
	s.queue2 = append(s.queue2, x)
	//将queue1中所有元素移到queue2中，再将两个队列进行交换
	s.Move()
}

func (s *MyStack) Move() {
	if len(s.queue1) == 0 {
		//交换，queue1置为queue2,queue2置为空
		s.queue1, s.queue2 = s.queue2, s.queue1
	} else {
		//queue1元素从头开始一个一个追加到queue2中
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:] //去除第一个元素
		s.Move()                //重复
	}
}

func (s *MyStack) Pop() int {
	val := s.queue1[0]
	s.queue1 = s.queue1[1:] //去除第一个元素
	return val
}

func (s *MyStack) Top() int {
	return s.queue1[0] //直接返回
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}
