package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	//从后序遍历中找到根节点
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootValue,
	}

	//如果没有叶子节点则直接返回
	if len(postorder) == 1 {
		return root
	}

	//在中序遍历中找到切割点
	delimiterIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootValue {
			delimiterIndex = i
			break
		}
	}

	//切割中序遍历数组(左闭右开)
	inorderLeft := inorder[0:delimiterIndex]
	inorderRight := inorder[delimiterIndex+1:]

	//根据中序遍历左右两部分的长度，切割后序遍历数组(左闭右开)
	postorderLeft := postorder[0:len(inorderLeft)]
	postorderRight := postorder[len(inorderLeft) : len(inorderLeft)+len(inorderRight)]

	//递归处理
	root.Left = buildTree(inorderLeft, postorderLeft)
	root.Right = buildTree(inorderRight, postorderRight)
	return root

}
