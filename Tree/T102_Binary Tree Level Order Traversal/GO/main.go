package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//用队列实现层序遍历
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	queue := list.New()
	queue.PushBack(root)

	tempArr := []int{}
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr = append(tempArr, node.Val)
		}
		ans = append(ans, tempArr)
		tempArr = []int{}
	}
	return ans
}

//递归法实现层序遍历
func levelOrder_Recursion(root *TreeNode) [][]int {
	ans := [][]int{}
	depth := 0
	var order func(*TreeNode, int)
	order = func(tn *TreeNode, deep int) {
		if tn == nil {
			return
		}
		if len(ans) == deep {
			ans = append(ans, []int{})
		}
		ans[deep] = append(ans[deep], tn.Val)
		order(tn.Left, deep+1)
		order(tn.Right, deep+1)
	}
	order(root, depth)
	return ans
}
