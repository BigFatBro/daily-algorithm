package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	// 双指针法，快慢指针
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}
