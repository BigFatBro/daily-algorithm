package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//迭代法
	dumyNode := &ListNode{}
	pre := dumyNode

	cur1, cur2 := list1, list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}

	if cur1 != nil {
		pre.Next = cur1
	} else {
		pre.Next = cur2
	}
	return dumyNode.Next
}
