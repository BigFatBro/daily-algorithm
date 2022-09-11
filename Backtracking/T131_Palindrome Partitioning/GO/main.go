package main

import "fmt"

func main() {

	fmt.Println("Output:", partition("aab"))
}

func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}

	var backtracking func(int)
	backtracking = func(startIndex int) {
		//终止条件
		if startIndex >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if isPalindrome(s, startIndex, i) {
				substr := s[startIndex : i+1]
				path = append(path, substr)

			} else {
				continue
			}
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return result

}

func isPalindrome(s string, start, end int) bool {
	i := start
	j := end
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true

}
