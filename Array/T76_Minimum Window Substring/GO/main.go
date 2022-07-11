package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("case1:", minWindow(s, t))

}

func minWindow(s string, t string) string {
	tCnt, sCnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		tCnt[t[i]] += 1
	}

	minLen := math.MaxInt32
	minLenL, minLenR := 0, 0

	check := func() bool {
		for k, v := range tCnt {
			if v > sCnt[k] {
				return false
			}
		}
		return true
	}

	start, end := 0, 0
	for end = 0; end < len(s); end++ {
		if tCnt[s[end]] > 0 {
			sCnt[s[end]]++
		}

		for check() && start <= end {
			if end-start+1 < minLen {
				minLen = end - start + 1
				minLenL = start
				minLenR = end + 1
			}
			// 移动窗口
			if _, ok := tCnt[s[start]]; ok {
				sCnt[s[start]] -= 1
			}

			start++
		}

	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLenL:minLenR]

}
