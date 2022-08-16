package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(*TreeNode)
	invert = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		tn.Left, tn.Right = tn.Right, tn.Left
		invert(tn.Left)
		invert(tn.Right)
	}

	invert(root)
	return root
}

func invertTree_Iter(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root

}
