package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return nil

}
