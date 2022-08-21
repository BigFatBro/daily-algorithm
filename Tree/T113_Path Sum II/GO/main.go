package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	curPath := []int{}
	var tranPathSum func(*TreeNode, int)
	tranPathSum = func(node *TreeNode, curSum int) {
		if node.Left == nil && node.Right == nil && curSum == targetSum {
			//tempPath := []int{}
			// tempPath = append(tempPath, curPath...)
			tempPath := make([]int, len(curPath))
			copy(tempPath, curPath)
			ans = append(ans, tempPath)
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}

		// 单层递归逻辑
		if node.Left != nil {
			curPath = append(curPath, node.Left.Val)
			tranPathSum(node.Left, curSum+node.Left.Val)
			curPath = curPath[:len(curPath)-1]
		}

		if node.Right != nil {
			curPath = append(curPath, node.Right.Val)
			tranPathSum(node.Right, curSum+node.Right.Val)
			curPath = curPath[:len(curPath)-1]
		}
		return
	}
	if root == nil {
		return ans
	}
	curPath = append(curPath, root.Val)
	tranPathSum(root, root.Val)
	return ans

}
