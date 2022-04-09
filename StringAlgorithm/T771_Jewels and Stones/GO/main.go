package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println("case1:", numJewelsInStones(jewels, stones))

}

func numJewelsInStones(jewels string, stones string) int {
	var jewelMap [128]bool
	for _, v := range jewels {
		jewelMap[v] = true
	}

	count := 0
	for _, v := range stones {
		if jewelMap[v] {
			count++
		}
	}

	return count
}
