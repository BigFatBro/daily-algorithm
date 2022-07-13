package main

type MyLinkedList struct {
	dummy *Node
	Size  int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
	}
	return MyLinkedList{dummy: rear, Size: 0}
}

func (this *MyLinkedList) Get(index int) int {

	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.dummy.Next
	curIndex := 0
	for {
		if curIndex == index {
			return cur.Val
		}
		curIndex++
		cur = cur.Next
	}

}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Val:  val,
		Next: this.dummy.Next,
	}

	this.dummy.Next = newNode
	this.Size++

}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		Val:  val,
		Next: nil,
	}

	cur := this.dummy
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = newNode
	this.Size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	} else if index < 0 {
		index = 0
	}

	newNode := &Node{
		Val: val,
	}
	cur := this.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || index >= this.Size {
		return
	}

	cur := this.dummy

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	this.Size--

}

func main() {

}
