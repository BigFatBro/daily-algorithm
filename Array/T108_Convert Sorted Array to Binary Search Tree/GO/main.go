package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	treeRoot := sortedArrayToBST(nums)
	fmt.Println("case1:", floorTraverse(treeRoot))
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := helper(nums, 0, len(nums)-1)
	return root
}

func helper(nums []int, left int, right int) *TreeNode {
	// 递归结束条件
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	// 创建一个节点
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

// 层序遍历（使用切片）
func floorTraverse(root *TreeNode) []int {
	ans := []int{}
	// 一个当作队列使用的切片
	temp := []TreeNode{}
	in, out := 0, 0

	// 根节点入队
	temp = append(temp, *root)
	in++

	for {
		if in <= out {
			break
		}
		// 出队
		ans = append(ans, temp[out].Val)
		if temp[out].Left != nil {
			temp = append(temp, *temp[out].Left)
			in++
		}
		if temp[out].Right != nil {
			temp = append(temp, *temp[out].Right)
			in++
		}
		out++
	}
	return ans
}
