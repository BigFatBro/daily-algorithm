package main

import "fmt"

func main() {
	case1_s := "egg"
	case1_t := "add"

	case2_s := "foo"
	case2_t := "bar"

	case3_s := "paper"
	case3_t := "title"

	fmt.Println("case1:", isIsomorphic(case1_s, case1_t))
	fmt.Println("case2:", isIsomorphic(case2_s, case2_t))
	fmt.Println("case3:", isIsomorphic(case3_s, case3_t))

	fmt.Println("case1:", isIsomorphicV2(case1_s, case1_t))
	fmt.Println("case2:", isIsomorphicV2(case2_s, case2_t))
	fmt.Println("case3:", isIsomorphicV2(case3_s, case3_t))

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLen := len(s)

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < sLen; i++ {
		sValue, sOk := sMap[s[i]]
		tValue, tOk := tMap[t[i]]

		if (sOk == true && sValue != t[i]) || (tOk == true && tValue != s[i]) {
			return false
		} else {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		}
	}
	return true
}

func isIsomorphicV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sLen := len(s)
	var sMap [128]int
	var tMap [128]int
	sPattern := ""
	tPattern := ""
	for i := 0; i < sLen; i++ {
		sAscii := int(s[i])
		tAscii := int(t[i])
		if sMap[sAscii] == 0 {
			sMap[sAscii] = i + 1
		}
		if tMap[tAscii] == 0 {
			tMap[tAscii] = i + 1
		}
		sPattern = sPattern + string(sMap[sAscii])
		tPattern = tPattern + string(tMap[tAscii])

	}
	if sPattern == tPattern {
		return true
	} else {
		return false
	}

}
