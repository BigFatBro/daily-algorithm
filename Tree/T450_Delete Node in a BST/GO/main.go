package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func deleteNode(root *TreeNode, key int) *TreeNode {
	//没找到待删除的节点
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			//找右子树的最左侧节点
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right

		}

	}
	return root
}

func deleteNode_Iter(root *TreeNode, key int) *TreeNode {

	var deleteOneNode func(*TreeNode) *TreeNode
	deleteOneNode = func(node *TreeNode) *TreeNode {
		if node == nil {
			return node
		} else if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left != nil && node.Right == nil {
			return node.Left
		} else if node.Right == nil && node.Left != nil {
			return node.Right
		} else {
			//找右子树的最左侧节点
			cur := node.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = node.Left
			return node.Right
		}
	}
	if root == nil {
		return root
	}

	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if pre == nil {
		//如果删除的是树的根节点
		return deleteOneNode(cur)
	}
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = deleteOneNode(cur)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = deleteOneNode(pre.Right)
	}
	return root

}
