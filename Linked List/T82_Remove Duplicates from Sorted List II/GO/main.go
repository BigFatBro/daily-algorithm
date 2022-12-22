package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	//先创建一个虚拟头节点
	dumyNode := &ListNode{
		Val: -1000,
	}
	dumyNode.Next = head
	//双指针
	cur := dumyNode

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			duplicateValue := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == duplicateValue {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dumyNode.Next
}
