package s232

type MyQueue struct {
	stackIn  []int //输入栈
	stackOut []int //输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func Exec() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	println(queue.Peek())
	println(queue.Pop())
	println(queue.Pop())
}

// Push
// @Description  往输入栈做push
// @Author lqc 2024-01-13 21:36:08
// @Param x
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// Pop
// @Description  在输出栈做pop，pop时如果输出栈数据为空，需要将输入栈全部数据导入，如果非空，则可直接使用
// @Author lqc 2024-01-13 21:36:21
// @Return int
func (this *MyQueue) Pop() int {
	inLen, outLen := len(this.stackIn), len(this.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}      //导出后清空
		outLen = len(this.stackOut) //更新长度值
	}
	val := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[:outLen-1]
	return val
}

// Peek
// @Description 返回队列最前面的值
// @Author lqc 2024-01-13 21:41:42
// @Return int
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

// Empty
// @Description 该队列是否为空
// @Author lqc 2024-01-13 21:41:56
// @Return bool
func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
