package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	//将链表的节点加入到切片中，使其可以用下标访问
	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	//Reorder链表
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
