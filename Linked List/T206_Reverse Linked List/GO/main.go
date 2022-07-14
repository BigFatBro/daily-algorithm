package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	dumyNode := &ListNode{}
	cur := dumyNode
	for _, v := range nums {
		temp := &ListNode{
			Val: v,
		}
		cur.Next = temp
		cur = cur.Next
	}

	showList(dumyNode.Next)

	head := reverseList(dumyNode.Next)
	showList(head)

}

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre

}

func showList(head *ListNode) {
	fmt.Print("\n")
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Print("nil")
}
