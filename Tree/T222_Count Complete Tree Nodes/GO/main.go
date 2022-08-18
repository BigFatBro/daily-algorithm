package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := root.Left
	right := root.Right
	leftHeight, rightHeight := 0, 0
	for left != nil {
		left = left.Left
		leftHeight++
	}
	for right != nil {
		right = right.Right
		rightHeight++
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1

}
