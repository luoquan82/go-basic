package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 5, 5, 4}

	current := nums[1]
	for _, n := range nums[1:] {
		current ^= n
	}

	fmt.Println("The number that appears only once is:", current)
}
