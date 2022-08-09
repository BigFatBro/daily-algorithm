package main

func main() {

}

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}

}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)

}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	if len(this.stackOut) == 0 {
		//输出栈为空，则把输入栈数据导入，再从输出栈弹出一个值
		for len(this.stackIn) != 0 {
			val := this.stackIn[len(this.stackIn)-1]
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
			this.stackOut = append(this.stackOut, val)
		}
	}

	val := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
