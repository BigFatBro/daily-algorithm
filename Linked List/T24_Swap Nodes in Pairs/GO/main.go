package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {

	// 创建一个虚拟头节点
	dummyNode := &ListNode{
		Next: head,
	}

	cur := dummyNode

	for cur.Next != nil && cur.Next.Next != nil {
		third := cur.Next.Next.Next

		first := cur.Next
		second := cur.Next.Next

		second.Next = first
		first.Next = third
		cur.Next = second

		cur = cur.Next.Next
	}

	return dummyNode.Next
}
