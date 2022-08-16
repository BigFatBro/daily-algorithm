package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

//递归法
func maxDepth(root *Node) int {
	var getDepth func(*Node) int
	getDepth = func(n *Node) int {
		if n == nil {
			return 0
		}

		//单层递归逻辑
		max := 0
		for _, v := range n.Children {
			depth := getDepth(v)
			if depth > max {
				max = depth
			}
		}
		return max + 1
	}

	return getDepth(root)
}

//迭代法，层序遍历
func maxDepth_Iter(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		depth++

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, v := range node.Children {
				if v != nil {
					queue = append(queue, v)
				}
			}
		}
	}
	return depth
}
