package main

import "fmt"

func main() {
	s := "25525511135"
	fmt.Println("Output:", restoreIpAddresses(s))

}

func restoreIpAddresses(s string) []string {
	result := []string{}
	path := []string{}
	var backtracking func(int, int)
	backtracking = func(startIndex, pointNum int) {
		if pointNum == 3 {
			if isValid(s, startIndex, len(s)-1) {
				result = append(result, path[0]+"."+path[1]+"."+path[2]+"."+s[startIndex:])

			}
			return
		}
		for i := startIndex; i < len(s); i++ {

			if isValid(s, startIndex, i) {
				path = append(path, s[startIndex:i+1])
				pointNum++
				backtracking(i+1, pointNum)
				pointNum--
				path = path[:len(path)-1]
			} else {
				break
			}

		}

	}

	if len(s) < 4 || len(s) > 12 {
		return result
	}
	backtracking(0, 0)
	return result

}

//判断字符串s在左闭又闭区间[start, end]所组成的数字是否合法
func isValid(s string, start, end int) bool {
	if start > end {
		return false
	}

	if s[start] == '0' && start != end {
		return false
	}
	num := 0
	for i := start; i <= end; i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}

		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}

	return true
}
