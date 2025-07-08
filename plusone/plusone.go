package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println("Plus One Example:", plusOne([]int{1, 2, 9})) // Output: [1, 2, 4]
}
