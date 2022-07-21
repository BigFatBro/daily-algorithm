package main

func main() {

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			mp[a+b]++
		}
	}
	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			if times, ok := mp[0-(c+d)]; ok {
				count += times
			}
		}
	}
	return count
}
