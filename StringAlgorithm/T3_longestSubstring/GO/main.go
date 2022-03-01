package main

import "fmt"

func main() {

	case1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(case1))
	case2 := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(case2))
	case3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(case3))

}

// lengthOfLongestSubstring用滑动窗口法寻找最长无重复字串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left := 0
	lookUp := make(map[byte]bool)
	n := len(s)
	maxLen := 0
	curLen := 0
	for i := 0; i < n; i++ {
		curLen++
		for {
			_, ok := lookUp[s[i]]
			if !ok {
				break
			}
			// 从Map中删除
			delete(lookUp, s[left])
			left++
			curLen--
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		lookUp[s[i]] = true
	}
	return maxLen

}
