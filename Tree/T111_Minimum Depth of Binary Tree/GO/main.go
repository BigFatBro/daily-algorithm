package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 递归法
func minDepth(root *TreeNode) int {
	var getDepth func(*TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)

		if node.Left == nil && node.Right != nil {
			return 1 + rightDepth
		} else if node.Left != nil && node.Right == nil {
			return 1 + leftDepth
		} else {
			if leftDepth < rightDepth {
				return leftDepth + 1
			} else {
				return rightDepth + 1
			}
		}

	}

	return getDepth(root)

}

//迭代法
func minDepth_Iter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 0

	for len(queue) > 0 {
		//一层的节点数量
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left == nil && node.Right == nil {
				return depth
			}

		}
	}

	return depth

}
