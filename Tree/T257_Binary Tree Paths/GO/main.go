package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func binaryTreePaths(root *TreeNode) []string {
	path := []int{}
	result := []string{}
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		path = append(path, node.Val)
		//递归终止逻辑
		if node.Left == nil && node.Right == nil {
			//遍历到叶节点，将路径保存
			sPath := ""
			for i := 0; i < len(path)-1; i++ {
				sPath = sPath + strconv.Itoa(path[i]) + "->"
			}
			sPath += strconv.Itoa(path[len(path)-1])
			result = append(result, sPath)
			return
		}

		// 单层递归逻辑
		if node.Left != nil {
			traversal(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			traversal(node.Right)
			path = path[:len(path)-1]
		}
	}
	traversal(root)
	return result

}

func binaryTreePaths_Iter(root *TreeNode) []string {
	nodeStack := []*TreeNode{}
	pathStack := []string{}
	result := []string{}
	if root == nil {
		return result
	}
	nodeStack = append(nodeStack, root)
	pathStack = append(pathStack, strconv.Itoa(root.Val))
	for len(nodeStack) != 0 {
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		path := pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]

		if node.Left == nil && node.Right == nil {
			result = append(result, path)
		}

		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(node.Right.Val))
		}
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(node.Left.Val))
		}
	}
	return result
}
