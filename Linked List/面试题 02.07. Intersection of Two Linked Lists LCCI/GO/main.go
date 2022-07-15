package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	for curB != nil {
		lenB++
		curB = curB.Next
	}

	curA = headA
	curB = headB
	if lenB > lenA {
		lenA, lenB = lenB, lenA
		curA, curB = curB, curA

	}

	//长度差
	gap := lenA - lenB

	//长链表的指针先移动
	for i := 0; i < gap; i++ {
		curA = curA.Next
	}

	//同时遍历长链表和短链表
	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
