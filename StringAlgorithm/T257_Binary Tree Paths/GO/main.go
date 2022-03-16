package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 1,
	}

	node2 := TreeNode{
		Val: 2,
	}

	node3 := TreeNode{
		Val: 3,
	}

	node5 := TreeNode{
		Val: 5,
	}

	root.Left = &node2
	root.Right = &node3
	node2.Right = &node5

	ans := binaryTreePaths(&root)
	fmt.Println(ans)
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	constructPaths(root, "", &paths)
	return paths
}

func constructPaths(root *TreeNode, path string, paths *[]string) {
	if root != nil {
		//path = path + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, path+strconv.Itoa(root.Val))
		} else {
			//path = path + "->"
			constructPaths(root.Left, path+strconv.Itoa(root.Val)+"->", paths)
			constructPaths(root.Right, path+strconv.Itoa(root.Val)+"->", paths)
		}
	}

}
