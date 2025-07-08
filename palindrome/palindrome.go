package main

import "fmt"

func isPalindrome(s string) bool {
	loop := len(s) / 2
	for i := range loop {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	p := "1234567654321"

	fmt.Println("The string is a palindrome:", isPalindrome(p))
}
