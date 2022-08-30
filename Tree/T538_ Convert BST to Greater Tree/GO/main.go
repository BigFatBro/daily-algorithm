package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func convertBST(root *TreeNode) *TreeNode {
	var pre int
	var traversal func(*TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Right)
		cur.Val += pre
		pre = cur.Val
		traversal(cur.Left)
	}

	pre = 0
	traversal(root)
	return root

}

//迭代法
func convertBST_Iter(root *TreeNode) *TreeNode {
	pre := 0
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Val += pre
			pre = node.Val

		}
	}
	return root
}
