package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumOfLeftLeaves(root *TreeNode) int {
	var getLeftLeaves func(*TreeNode) int
	getLeftLeaves = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftValue := getLeftLeaves(node.Left)
		rightValue := getLeftLeaves(node.Right)

		midValue := 0
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			midValue = node.Left.Val
		}
		return leftValue + rightValue + midValue
	}
	result := getLeftLeaves(root)
	return result

}

//迭代法
func sumOfLeftLeaves_Iter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	result := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	return result
}
