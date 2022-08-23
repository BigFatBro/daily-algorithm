package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// rootValue := 0
	// if root1 == nil && root2 == nil {
	// 	return nil
	// } else if root1 != nil && root2 == nil {
	// 	rootValue = root1.Val
	// } else if root1 == nil && root2 != nil {
	// 	rootValue = root2.Val
	// } else {
	// 	rootValue = root1.Val + root2.Val
	// }
	root := &TreeNode{
		Val: root1.Val + root2.Val,
	}

	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root

}

func mergeTrees_Iter(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	que := []*TreeNode{}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	que = append(que, root1, root2)
	for len(que) > 0 {
		node1 := que[0]
		que = que[1:]
		node2 := que[0]
		que = que[1:]
		node1.Val += node2.Val

		if node1.Left != nil && node2.Left != nil {
			que = append(que, node1.Left, node2.Left)
		}

		if node1.Right != nil && node2.Right != nil {
			que = append(que, node1.Right, node2.Right)
		}

		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}

		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}

	}

	return root1

}
