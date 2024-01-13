package s239

// 封装单调队列的方式解题
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	//  当队列不为空且该值比队列尾巴的值大
	//  把该尾巴制空，长度减一
	//  直到插入的值小于队列尾巴
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}

	//  正常添加
	//  队列为空或者该值比队列尾巴小
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

// MaxSlidingWindow
// @Description 滑动窗口最大值
// @Author lqc 2024-01-13 22:31:20
// @Param nums
// @Param k
// @Return []int
func MaxSlidingWindow(nums []int, size int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < size; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := size; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-size])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}
