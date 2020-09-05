package main

import "fmt"

func main() {

	fmt.Println("Is palindrome", palindrome("MADAMFDE"))

	fmt.Println("Is palindrome", palindrome("MADAM"))

}

func palindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
