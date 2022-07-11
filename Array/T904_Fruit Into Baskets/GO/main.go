package main

import "fmt"

func main() {
	case1 := []int{1, 2, 1}
	case2 := []int{0, 1, 2, 2}
	case3 := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println("case1:", totalFruit(case1))
	fmt.Println("case2:", totalFruit(case2))
	fmt.Println("case2:", totalFruit(case3))

}

func totalFruit(fruits []int) int {
	//滑动窗口法
	maxFruitNum := 0
	start, end := 0, 0
	categoryCounter := map[int]int{}

	for end = 0; end < len(fruits); end++ {
		categoryCounter[fruits[end]] += 1

		for len(categoryCounter) >= 3 {
			categoryCounter[fruits[start]] -= 1
			if categoryCounter[fruits[start]] <= 0 {
				delete(categoryCounter, fruits[start])
			}
			start++
		}
		if end-start+1 > maxFruitNum {
			maxFruitNum = end - start + 1
		}
	}

	return maxFruitNum

}
