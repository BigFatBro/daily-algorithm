package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		right := trimBST(root.Right, low, high)
		return right
	}

	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root

}

func trimBST_Iter(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//处理头节点,将头节点移动到[low,high]
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}

	}
	//处理root左子树中小于low的节点
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}

	// 处理root右子树中大于high的节点
	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root

}
