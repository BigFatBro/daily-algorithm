package main

import "sort"

func main() {

}

func groupAnagramsV1(strs []string) [][]string {
	mp := map[string][]string{}
	//方法1：将字符串按字母排序后组成的字符串作为map的key
	for _, v := range strs {
		chars := []byte(v)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		sortedStr := string(chars)
		mp[sortedStr] = append(mp[sortedStr], v)
	}

	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagramsV2(strs []string) [][]string {
	// 方法2：统计每个字符串各字母出现的个数，把这个统计结果作为map的key
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cnt[v-'a']++
		}

		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
