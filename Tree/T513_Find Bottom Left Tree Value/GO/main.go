package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findBottomLeftValue(root *TreeNode) int {
	var mostLeftValue int
	maxDeepth := -1

	var getMostLeftValue func(*TreeNode, int)
	getMostLeftValue = func(node *TreeNode, depth int) {
		//递归终止条件
		if node.Left == nil && node.Right == nil {
			if depth > maxDeepth {
				maxDeepth = depth
				mostLeftValue = node.Val
			}
			return
		}
		//单层递归逻辑
		if node.Left != nil {
			getMostLeftValue(node.Left, depth+1) //隐藏着回溯
		}
		if node.Right != nil {
			getMostLeftValue(node.Right, depth+1)
		}
		return

	}
	getMostLeftValue(root, 0)
	return mostLeftValue

}

//迭代法：层序遍历
func findBottomLeftValue_Iter(root *TreeNode) int {
	var ans int
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans

}
