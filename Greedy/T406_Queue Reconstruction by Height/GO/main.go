package main

import "sort"

func main() {

}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		//身高相同时按照k从小到大排序
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			//身高不同时，按照身高从大到小排序
			return people[i][0] > people[j][0]
		}
	})

	//再按照k插入
	result := make([][]int, 0)
	for _, info := range people {
		//先添加到末尾
		result = append(result, info)
		//从插入的目标位置向后移动
		copy(result[info[1]+1:], result[info[1]:])
		//把末尾的元素添加到目标位置
		result[info[1]] = info
	}
	return result

}
