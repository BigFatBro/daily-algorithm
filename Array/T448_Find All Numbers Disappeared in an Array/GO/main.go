package main

import (
	"fmt"
	"sort"
)

func main() {
	case1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	case2 := []int{1, 1}
	fmt.Println("case1:", findDisappearedNumbers(case1))
	fmt.Println("case1V2:", findDisappearedNumbersV2(case1))
	fmt.Println("case2V2:", findDisappearedNumbersV2(case2))
	fmt.Println("case1V3:", findDisappearedNumbersV3(case1))

}

// 方法1：哈希表
func findDisappearedNumbers(nums []int) []int {
	ans := []int{}
	table := make(map[int]bool)
	for _, v := range nums {
		table[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := table[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}

//方法2(垃圾解法，击败7.34%的用户)：先排序，再遍历
func findDisappearedNumbersV2(nums []int) []int {
	ans := []int{}
	sort.Ints(nums)
	target := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			target++
		} else {
			if nums[i] > target {
				ans = append(ans, target)
				target++
				i--
				continue
			}
		}
	}
	// 处理末尾情况
	if nums[len(nums)-1] != target && target <= len(nums) {
		for i := target; i <= len(nums); i++ {
			ans = append(ans, i)
		}
	}

	return ans
}

//方法3：对方法1优化，不使用额外空间
func findDisappearedNumbersV3(nums []int) []int {
	for _, v := range nums {
		//将nums[v-1]置为负数，因为nums[v-1]可能已经被其他人置为负数了，所以每次使用v之前要先取绝对值
		nums[abs(v)-1] = -abs(nums[abs(v)-1])
	}

	ans := []int{}
	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
