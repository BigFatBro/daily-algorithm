package main

import "sort"

func main() {

}

func findContentChildren(g []int, s []int) int {
	//贪心算法
	sort.Ints(s)
	sort.Ints(g)
	index := len(s) - 1
	res := 0
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			res++
			index--
		}
	}
	return res

}
