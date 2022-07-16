package main

import "fmt"

func main() {
	s := "abab"
	p := "ab"
	fmt.Println("case1:", findAnagrams(s, p))

}

func findAnagrams(s string, p string) []int {
	ans := []int{}
	cnt := [26]int{}
	for _, v := range p {
		cnt[v-'a']++
	}

	sLen := len(s)
	pLen := len(p)

	for i := 0; i <= sLen-pLen; i++ {
		curCnt := [26]int{}
		for j := 0; j < pLen; j++ {
			curCnt[s[i+j]-'a']++
		}
		if curCnt == cnt {
			ans = append(ans, i)
		}

	}

	return ans

}

func findAnagramsV2(s string, p string) []int {
	// 滑动窗口法
	ans := []int{}
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	var sCnt, pCnt [26]int
	for i, v := range p {
		sCnt[s[i]-'a']++
		pCnt[v-'a']++
	}

	if sCnt == pCnt {
		ans = append(ans, 0)
	}

	//开始滑动窗口
	for i, v := range s[:sLen-pLen] {
		sCnt[v-'a']--
		sCnt[s[i+pLen]-'a']++
		if sCnt == pCnt {
			ans = append(ans, i+1)
		}
	}

	return ans
}
