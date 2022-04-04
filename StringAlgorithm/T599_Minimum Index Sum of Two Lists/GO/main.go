package main

import (
	"fmt"
	"math"
)

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	ans := findRestaurant(list1, list2)
	fmt.Println("ans:", ans)

}

func findRestaurant(list1 []string, list2 []string) []string {
	l1Map := make(map[string]int, len(list1))
	for i, v := range list1 {
		l1Map[v] = i
	}
	minIndexSum := math.MaxInt32
	ans := []string{}
	for i, v := range list2 {
		if index, ok := l1Map[v]; ok {
			indexSum := i + index
			if indexSum == minIndexSum {
				ans = append(ans, v)
			} else if indexSum < minIndexSum {
				minIndexSum = indexSum
				ans = []string{v}
			}
		}
	}
	return ans
}
