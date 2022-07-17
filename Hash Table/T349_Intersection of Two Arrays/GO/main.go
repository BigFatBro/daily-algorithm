package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("case1:", intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	for _, v := range nums1 {
		mp[v] = 1
	}
	ans := []int{}
	for _, v := range nums2 {
		if count, ok := mp[v]; ok && count > 0 {
			ans = append(ans, v)
			mp[v]--
		}
	}

	return ans

}
