package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	//从前序遍历中取出根节点
	rootValue := preorder[0]
	root := &TreeNode{
		Val: rootValue,
	}

	if len(preorder) == 1 {
		return root
	}

	//找出中序遍历的分割点
	delimiterIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootValue {
			delimiterIndex = i
			break
		}
	}

	//切割中序遍历数组
	inorderLeft := inorder[0:delimiterIndex]
	inorderRight := inorder[delimiterIndex+1:]

	//根据中序遍历切出来的两部分数组的长度，切割先序遍历数组
	preorderLeft := preorder[1 : 1+len(inorderLeft)]
	preorderRight := preorder[1+len(inorderLeft) : 1+len(inorderLeft)+len(inorderRight)]

	//递归处理
	root.Left = buildTree(preorderLeft, inorderLeft)
	root.Right = buildTree(preorderRight, inorderRight)
	return root
}
