package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "5F3Z-2e-9-w"
	k1 := 4
	s2 := "2-5g-3-J"
	k2 := 2

	fmt.Println("case1:", licenseKeyFormatting(s1, k1))
	fmt.Println("case2:", licenseKeyFormatting(s2, k2))

}

func licenseKeyFormatting(s string, k int) string {
	sLen := len(s)
	alphanumerics := []byte{}

	for i := 0; i < sLen; i++ {
		if s[i] != '-' {
			alphanumerics = append(alphanumerics, s[i])
		}
	}

	ans := ""
	alphanumericsString := strings.ToUpper(string(alphanumerics))
	alphanumericsLen := len(alphanumericsString)

	if alphanumericsLen <= k {
		return alphanumericsString
	}

	i := alphanumericsLen - k

	for ; i >= 0; i = i - k {
		if ans != "" {
			ans = alphanumericsString[i:i+k] + "-" + ans
		} else {
			ans = alphanumericsString[i:i+k] + ans
		}

	}
	if i > 0-k {
		ans = alphanumericsString[0:i+k] + "-" + ans
	}

	return ans

}

func licenseKeyFormattingV2(s string, k int) string {
	ans := []byte{}
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if s[i] != '-' {
			ans = append(ans, byte(unicode.ToUpper(rune(s[i]))))
			cnt++
			if cnt%k == 0 {
				ans = append(ans, '-')
			}
		}

	}

	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
