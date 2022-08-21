package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var haspathsum func(*TreeNode, int) bool
	haspathsum = func(node *TreeNode, cnt int) bool {
		if node.Left == nil && node.Right == nil && cnt == 0 {
			return true
		}
		if node.Left == nil && node.Right == nil && cnt != 0 {
			return false
		}
		//单层递归逻辑
		if node.Left != nil {
			if haspathsum(node.Left, cnt-node.Left.Val) {
				return true
			}
		}

		if node.Right != nil {
			if haspathsum(node.Right, cnt-node.Right.Val) {
				return true
			}
		}
		return false
	}
	if root == nil {
		return false
	}

	return haspathsum(root, targetSum-root.Val)

}

func hasPathSum_Iter(root *TreeNode, targetSum int) bool {
	nodeStack := []*TreeNode{}
	valueSumStack := []int{}
	if root == nil {
		return false
	}
	nodeStack = append(nodeStack, root)
	valueSumStack = append(valueSumStack, root.Val)
	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		valueSum := valueSumStack[len(valueSumStack)-1]
		valueSumStack = valueSumStack[:len(valueSumStack)-1]
		if valueSum == targetSum {
			return true
		}
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			valueSumStack = append(valueSumStack, valueSum+node.Left.Val)
		}
		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			valueSumStack = append(valueSumStack, valueSum+node.Right.Val)
		}
	}
	return false
}
