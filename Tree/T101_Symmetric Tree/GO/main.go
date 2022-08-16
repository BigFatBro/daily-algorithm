package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right != nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		// 左右节点不为空，数值相同时应该继续递归，而不应该返回
		//单层递归
		// 处理外侧
		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		//二者都为true时返回true,否则返回false
		return outside && inside
	}
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)

}

func isSymmetric_Iter(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		//出栈
		left := queue[0]
		queue = queue[1:]
		right := queue[0]
		queue = queue[1:]

		if left == nil && right == nil {
			//左右节点为空，依旧对称，继续下一次循环
			continue
		}
		if left == nil || right == nil || (left.Val != right.Val) {
			return false
		}

		//入栈其子节点
		//外侧
		queue = append(queue, left.Left, right.Right)
		//内侧
		queue = append(queue, left.Right, right.Left)

	}
	return true

}
