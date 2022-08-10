package main

import "fmt"

func main() {

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println("Output:", maxSlidingWindow(nums, k))

}

//单调队列
type MyQueue struct {
	q []int
}

func Constructor() MyQueue {
	return MyQueue{
		q: make([]int, 0),
	}
}

func (this *MyQueue) pop(value int) {
	if len(this.q) != 0 && value == this.q[0] {
		this.q = this.q[1:]
	}

}

func (this *MyQueue) push(value int) {
	for len(this.q) != 0 && value > this.q[len(this.q)-1] {
		this.q = this.q[:len(this.q)-1]
	}

	this.q = append(this.q, value)

}

func (this *MyQueue) front() int {
	return this.q[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	myQ := Constructor()
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		myQ.push(nums[i])
	}
	ans = append(ans, myQ.front())
	for i := k; i < len(nums); i++ {
		myQ.pop(nums[i-k])
		myQ.push(nums[i])
		ans = append(ans, myQ.front())
	}
	return ans

}
