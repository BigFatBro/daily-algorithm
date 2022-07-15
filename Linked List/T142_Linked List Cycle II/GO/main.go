package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		//慢指针每次走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			//分别从head和相遇点同时遍历
			index1 := fast
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index2
		}
	}
	return nil
}
