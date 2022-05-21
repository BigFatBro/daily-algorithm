package main

import (
	"fmt"
	"strconv"
)

func main() {

	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("Output1:", summaryRanges(nums))

}

// 方法1：双指针法,一次遍历
func summaryRanges(nums []int) []string {
	ans := []string{}
	var i int
	var j int
	for i = 0; i < len(nums); {
		j = i + 1
		for {
			if j >= len(nums) || nums[j] != nums[j-1]+1 {
				break
			}
			j++
		}

		if j-1 == i {
			ans = append(ans, strconv.Itoa(nums[i]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
		}
		i = j
	}
	return ans

}
