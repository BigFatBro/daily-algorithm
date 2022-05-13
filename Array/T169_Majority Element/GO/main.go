package main

import "fmt"

func main() {
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	nums2 := []int{3, 2, 3}
	nums3 := []int{3, 2, 2}
	nums4 := []int{1, 3, 1, 1, 4, 1, 1, 5, 1, 1, 6, 2, 2}
	fmt.Println("Output1:", majorityElement(nums1))
	fmt.Println("Output2:", majorityElement(nums2))
	fmt.Println("Output3:", majorityElement(nums3))
	fmt.Println("Output4:", majorityElement(nums4))

}

func majorityElement(nums []int) int {
	// majority element is the element that appears more than ⌊n / 2⌋ times
	// 使用守阵地思想,也叫投票法
	majEle := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majEle = nums[i]
			count = 1
			continue
		}
		if nums[i] == majEle {
			count++
		} else {
			count--
		}
	}
	return majEle
}
