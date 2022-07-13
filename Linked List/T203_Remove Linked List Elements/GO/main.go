package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	var head = new(ListNode)
	var tail *ListNode
	tail = head
	for _, v := range nums {
		temp := &ListNode{
			Val: v,
		}
		tail.Next = temp
		tail = temp
	}

	head = head.Next
	showList(head)

	fmt.Println("\ndelete val: 6")
	head = removeElements(head, 6)
	showList(head)

}

func removeElements(head *ListNode, val int) *ListNode {
	//处理头节点
	for head != nil && head.Val == val {
		head = head.Next
	}

	//处理非头节点
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return head

	//也可以直接创建一个虚拟头节点，这样就不用特殊处理头节点

}

func showList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print("->", cur.Val)
		cur = cur.Next
	}
}
