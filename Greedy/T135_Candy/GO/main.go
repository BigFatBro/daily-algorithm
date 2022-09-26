package main

import "fmt"

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println("Output:", candy(ratings))
}

func candy(ratings []int) int {

	candyDis := make([]int, len(ratings))
	for i := 0; i < len(candyDis); i++ {
		candyDis[i] = 1
	}

	//从前向后
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyDis[i] = candyDis[i-1] + 1
		}
	}

	//从后向前
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candyDis[i] < (candyDis[i+1] + 1) {
				candyDis[i] = candyDis[i+1] + 1
			}

		}
	}

	result := 0
	for _, v := range candyDis {
		result += v
	}

	return result

}
