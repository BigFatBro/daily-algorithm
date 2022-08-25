package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var travel func(*TreeNode, *TreeNode, *TreeNode) *TreeNode
	travel = func(cur, p, q *TreeNode) *TreeNode {
		if cur == nil {
			return cur
		}

		if cur.Val > p.Val && cur.Val > q.Val {
			left := travel(cur.Left, p, q)
			if left != nil {
				return left
			}
		}

		if cur.Val < p.Val && cur.Val < q.Val {
			right := travel(cur.Right, p, q)
			if right != nil {
				return right
			}
		}

		return cur
	}

	return travel(root, p, q)
}

//迭代法
func lowestCommonAncestor_Iter(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
