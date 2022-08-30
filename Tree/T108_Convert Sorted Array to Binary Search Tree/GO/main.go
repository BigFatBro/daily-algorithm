package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	var traversal func(int, int) *TreeNode
	traversal = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = traversal(left, mid-1)
		root.Right = traversal(mid+1, right)
		return root
	}

	return traversal(0, len(nums)-1)
}
