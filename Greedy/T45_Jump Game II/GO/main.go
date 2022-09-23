package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println("Output:", jump(nums))

}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	curDistance, netxDistance := 0, 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > netxDistance {
			netxDistance = nums[i] + i
		}

		if i == curDistance {
			if curDistance != len(nums)-1 {
				ans++
				curDistance = netxDistance
				if curDistance >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}

	}
	return ans
}
