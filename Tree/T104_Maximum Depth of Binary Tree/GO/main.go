package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//递归法
func maxDepth(root *TreeNode) int {
	var getDepth func(*TreeNode) int
	getDepth = func(node *TreeNode) int {
		//递归终止条件
		if node == nil {
			return 0
		}

		//单层递归逻辑
		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}

	return getDepth(root)
}

//迭代法，层序遍历
func maxDepth_Iter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		depth++
		//每次循环出队，要出树的一层的全部节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth

}
