package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//双指针
	p1, p2 := l1, l2
	//创建一个新的链表的虚拟头节点
	dumyNode := &ListNode{}
	p := dumyNode

	//加法进位
	c := 0
	newValue := 0

	for p1 != nil || p2 != nil {
		sum := 0
		if p1 != nil && p2 != nil {
			sum = p1.Val + p2.Val + c
		} else if p1 != nil && p2 == nil {
			sum = p1.Val + c
		} else if p2 != nil && p1 == nil {
			sum = p2.Val + c
		}

		if sum >= 10 {
			newValue = (sum % 10)
			c = sum / 10
		} else {
			newValue = sum
			c = 0
		}
		//创建一个新节点
		p.Next = &ListNode{
			Val: newValue,
		}

		p = p.Next

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}

	}

	// 处理加法进位问题
	if c != 0 {
		p.Next = &ListNode{
			Val: c,
		}
		p = p.Next
	}

	return dumyNode.Next
}
