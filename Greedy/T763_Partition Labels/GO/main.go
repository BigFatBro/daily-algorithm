package main

import "fmt"

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println("Output:", partitionLabels(S))
}

func partitionLabels(s string) []int {
	hash := make([]int, 27)
	for i := 0; i < len(s); i++ {
		//统计每个字符最后出现的位置
		hash[s[i]-'a'] = i
	}
	result := []int{}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if hash[s[i]-'a'] > right {
			right = hash[s[i]-'a']
		}
		if i == right {
			result = append(result, i-left+1)
			left = i + 1
		}
	}
	return result

}
