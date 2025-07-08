package main

import "fmt"

func removeDuplicate(nums []int) int {
	pi, pk := 0, 1
	for pk < len(nums) {
		if nums[pi] != nums[pk] {
			pi++
			nums[pi] = nums[pk]
		}
		pk++
	}
	return pi + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4}
	lenth := removeDuplicate(nums)
	fmt.Println(lenth)
}
