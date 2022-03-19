package main

import "fmt"

func main() {
	case1_ransomNote := "aa"
	case1_magazine := "ab"
	fmt.Println("case1:", canConstruct(case1_ransomNote, case1_magazine))
	case2_ransomNote := "aa"
	case2_magazine := "aab"
	fmt.Println("case1:", canConstruct(case2_ransomNote, case2_magazine))

}

func canConstruct(ransomNote string, magazine string) bool {
	rMap := make(map[byte]int)
	mMap := make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		rMap[ransomNote[i]] = rMap[ransomNote[i]] + 1
	}
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]] = mMap[magazine[i]] + 1
	}

	for k, v := range rMap {
		count, ok := mMap[k]
		if !ok || (count-v) < 0 {
			return false
		}
	}
	return true

}

func canConstruct2(ransomNote string, magazine string) bool {
	//因为ransomNote和magazine由a~z之间的字母组成，所以上述算法可以用数组优化
	if len(magazine) < len(ransomNote) {
		return false
	}
	cnt := [26]int{}
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}

	return true

}
