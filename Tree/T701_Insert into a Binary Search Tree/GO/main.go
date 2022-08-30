package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{
			Val: val,
		}
		return node
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST_Iter(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	cur := root
	parent := cur
	for cur != nil {
		parent = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	node := &TreeNode{
		Val: val,
	}
	if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return root
}
