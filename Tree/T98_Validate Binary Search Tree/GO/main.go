package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 验证中序遍历序列是否是单调递增即可
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var travel func(*TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftRes := travel(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}

	return travel(root)

}

// 迭代法
func isValidBST_Iter(root *TreeNode) bool {
	stack := []*TreeNode{}
	var pre *TreeNode
	if root != nil {
		stack = append(stack, root)
	} else {
		return true
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && node.Val <= pre.Val {
				return false
			}
			pre = node

		}
	}
	return true
}
