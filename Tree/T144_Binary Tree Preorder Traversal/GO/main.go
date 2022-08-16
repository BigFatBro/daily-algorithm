package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//先序遍历递归实现
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var preorder func(*TreeNode)
	preorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}

		ans = append(ans, tn.Val)
		preorder(tn.Left)
		preorder(tn.Right)
	}
	preorder(root)
	return ans
}

//先序遍历迭代实现
func preorderTraversal_Iter(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		//出栈一个节点，并将其右孩子入栈
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans

}

func preorderTraversal_Iter_v2(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return ans
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		//出栈一个节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//将其右孩子和左孩子入栈,先右后左
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

//后序遍历递归实现
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var postorder func(*TreeNode)
	postorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		postorder(tn.Left)
		postorder(tn.Right)
		ans = append(ans, tn.Val)
	}
	postorder(root)
	return ans
}

//后序遍历迭代实现
//先序是中左右，后序是左右中
//将先序变为中右左，在反转为左右中，即可得到后续遍历结果
func postorderTraversal_Iter(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return ans
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Left.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	//反转遍历结果
	for i, j := 0, len(ans)-1; j > i; i++ {
		ans[i], ans[j] = ans[j], ans[i]
		j--

	}
	return ans

}

//中序遍历递归实现
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		ans = append(ans, tn.Val)
		inorder(tn.Right)

	}
	inorder(root)
	return ans
}

//中序遍历迭代实现
func inorderTraversal_Iter(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return ans
	}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}
