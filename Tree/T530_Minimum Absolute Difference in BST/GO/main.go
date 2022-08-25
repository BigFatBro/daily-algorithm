package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var pre *TreeNode
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)
		if pre != nil && result < (node.Val-pre.Val) {
			result = (node.Val - pre.Val)
		}
		pre = node

		travel(node.Right)
	}
	travel(root)
	return result

}
