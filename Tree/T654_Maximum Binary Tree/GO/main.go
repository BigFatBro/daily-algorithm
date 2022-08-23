package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	node := &TreeNode{}
	//终止条件
	if len(nums) == 1 {
		node.Val = nums[0]
	}

	//单层递归逻辑
	//寻找最大值
	maxNum := math.MinInt
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	node.Val = maxNum
	if maxIndex > 0 {
		node.Left = constructMaximumBinaryTree(nums[:maxIndex])
	}

	if maxIndex < len(nums)-1 {
		node.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	}

	return node

	//优化方法：不新建切片，递归时传入索引

}
