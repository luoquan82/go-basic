package main

import "fmt"

func targetSum(nums []int, target int) map[int]int {
	result := make(map[int]int)
	for i, v := range nums {
		for _, w := range nums[i+1:] {
			if v+w == target {
				result[v] = w // Store the indices of the two numbers
			}
		}
	}
	return result
}

func main() {
	target := 12
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := targetSum(nums, target)
	fmt.Println("Pairs that sum to", target, ":", ret)
}
