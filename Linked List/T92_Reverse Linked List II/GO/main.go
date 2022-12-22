package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dumyNode := &ListNode{}
	dumyNode.Next = head

	preSectionLastNode := dumyNode
	for i := 1; i < left; i++ {
		preSectionLastNode = preSectionLastNode.Next
	}
	sectionStart := preSectionLastNode.Next
	sectionEnd := sectionStart
	for i := 0; i < right-left; i++ {
		sectionEnd = sectionEnd.Next
	}
	postSectionfirstNode := sectionEnd.Next

	//断开链表
	preSectionLastNode.Next = nil
	sectionEnd.Next = nil

	//反转以sectionStart开始的链表
	var temp, cur, pre *ListNode
	pre = postSectionfirstNode
	cur = sectionStart
	for cur != nil {
		temp = cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	//连接完成反转后的链表
	preSectionLastNode.Next = pre
	return dumyNode.Next

}
