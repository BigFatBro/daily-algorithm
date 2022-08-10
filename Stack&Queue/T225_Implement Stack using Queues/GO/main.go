package main

func main() {

}

type MyStack struct {
	que1 []int
	que2 []int
}

func Constructor() MyStack {
	return MyStack{
		que1: make([]int, 0),
		que2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.que1 = append(this.que1, x)
}

func (this *MyStack) Pop() int {
	if len(this.que1) == 0 {
		return 0
	}

	for len(this.que1) > 1 {
		this.que2 = append(this.que2, this.que1[0])
		this.que1 = this.que1[1:]
	}
	result := this.que1[0]

	this.que1 = []int{}
	this.que1 = append(this.que1, this.que2...)
	this.que2 = []int{}
	return result
}

func (this *MyStack) Top() int {
	return this.que1[len(this.que1)-1]

}

func (this *MyStack) Empty() bool {
	if len(this.que1) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
