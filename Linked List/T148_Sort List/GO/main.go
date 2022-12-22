package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func sortList(head *ListNode) *ListNode {
	//归并排序（递归）
	if head == nil || head.Next == nil {
		return head
	}
	//分割
	//快慢指针找到链表的中间位置
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	// mid为分割后的后面一段的第一个节点
	mid := slow.Next
	// 截断链表
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	//归并
	//先建立一个虚拟头节点
	dumyNode := &ListNode{}
	h := dumyNode
	for left != nil && right != nil {
		if left.Val <= right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}

	return dumyNode.Next

}
