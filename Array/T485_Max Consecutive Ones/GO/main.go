package main

import "fmt"

func main() {
	case1 := []int{1, 1, 0, 1, 1, 1}
	fmt.Println("case1:", findMaxConsecutiveOnes(case1))
	case2 := []int{1, 0, 1, 1, 0, 1}
	fmt.Println("case1:", findMaxConsecutiveOnes(case2))
}

func findMaxConsecutiveOnes(nums []int) int {
	maxConsOnes := 0
	numsLen := len(nums)
	var i, j int

	for i = 0; i < numsLen; {
		if nums[i] == 1 {
			j = i + 1
			curOnes := 1
			for {
				if j < numsLen && nums[j] == 1 {
					curOnes++
					j++
				} else {
					break
				}
			}
			if curOnes > maxConsOnes {
				maxConsOnes = curOnes
			}
			i = i + curOnes
		} else {
			i++
		}
	}
	return maxConsOnes

}

// 官方解法表达的更简洁
func findMaxConsecutiveOnesV2(nums []int) (maxCnt int) {
	cnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}
	maxCnt = max(maxCnt, cnt)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
