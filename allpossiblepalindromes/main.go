package main

import "fmt"

func main() {
	s := "aaacac"

	allPossiblePalindromes(s)
}

func allPossiblePalindromes(s string) {

	for i, inp := range s {
		fmt.Println("palindrome is ", string(inp))
		test := 1
		len := len(s) - i - 1
		var testString string
		testString = string(s[i])
		for test <= len {
			testString = testString + string(s[test+i])
			if isPalindrome(testString) {
				fmt.Println("Palindrome is ", testString)
			}
			test++
		}
	}
}

func isPalindrome(n string) bool {

	if len(n) == 0 {
		return true
	}

	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-i-1] {
			return false
		}
	}
	return true
}
