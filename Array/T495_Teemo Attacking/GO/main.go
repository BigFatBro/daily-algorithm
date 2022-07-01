package main

import "fmt"

func main() {
	timeSeries := []int{1, 2}
	duration := 2
	fmt.Println("case1:", findPoisonedDuration(timeSeries, duration))

}

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	expired := 0
	for _, v := range timeSeries {
		if v >= expired {
			ans += duration
		} else {
			ans += v + duration - expired
		}
		expired = v + duration
	}
	return ans

}
