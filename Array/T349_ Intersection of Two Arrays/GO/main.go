package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("Output:", intersection(nums1, nums2))
	fmt.Println("OutputV2:", intersectionV2(nums1, nums2))

}

//方法1：哈希表（因为本题数组中的数字是0~1000，所以可以直接用数组当哈希表）
func intersection(nums1 []int, nums2 []int) []int {
	table := make([]int, 1001)
	ans := []int{}

	for _, v := range nums1 {
		table[v] = 1
	}

	for _, v := range nums2 {
		if table[v] == 1 {
			ans = append(ans, v)
			table[v]++

		}
	}
	return ans
}

//方法2：排序加双指针法
func intersectionV2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := []int{}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(ans) == 0 || x > ans[len(ans)-1] {
				ans = append(ans, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return ans

}
