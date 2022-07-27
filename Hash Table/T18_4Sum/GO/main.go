package main

import (
	"sort"
)

func main() {

}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	if n < 4 {
		return ans
	}

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if (nums[i] + nums[i+1] + nums[i+2] + nums[i+3]) > target {
			break
		}
		if (nums[i] + nums[n-3] + nums[n-2] + nums[n-1]) < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if (nums[i] + nums[j] + nums[j+1] + nums[j+2]) > target {
				break
			}
			if (nums[i] + nums[j] + nums[n-2] + nums[n-1]) < target {
				continue
			}

			left := j + 1
			right := n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++

					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}

			}

		}
	}
	return ans
}
