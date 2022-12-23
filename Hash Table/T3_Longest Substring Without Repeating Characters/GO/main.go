package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println("Output:", lengthOfLongestSubstring(s))

}

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口法
	//哈希集合，记录窗口内的每个字符是否出现过
	m := map[byte]int{}
	n := len(s)

	//窗口最大长度
	ans := 0

	//窗口结束指针
	rk := -1

	for i := 0; i < n; i++ {
		if i != 0 {
			// 窗口起始指针， 没移动一格移除一个字符
			delete(m, s[i-1])
		}

		// 不断移动窗口结束指针
		for rk+1 < n && m[s[rk+1]] == 0 {
			//记录字符出现过
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}

}
