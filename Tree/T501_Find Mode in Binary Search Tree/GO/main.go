package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findMode(root *TreeNode) []int {
	var pre *TreeNode
	maxCount := 0
	count := 0
	result := []int{}
	var travel func(*TreeNode)
	travel = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		travel(cur.Left)
		//中序遍历
		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		pre = cur
		if count == maxCount {
			result = append(result, cur.Val)
		}
		if count > maxCount {
			maxCount = count
			result = []int{}
			result = append(result, cur.Val)
		}
		travel(cur.Right)
	}
	travel(root)
	return result

}
