package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := getHeight(node.Left)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := getHeight(node.Right)
		if rightHeight == -1 {
			return -1
		}

		if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
			return -1
		}

		if leftHeight > rightHeight {
			return 1 + leftHeight
		} else {
			return 1 + rightHeight
		}

	}

	result := getHeight(root)
	if result == -1 {
		return false
	} else {
		return true
	}

}

func isBalanced_Iter(root *TreeNode) bool {

}
