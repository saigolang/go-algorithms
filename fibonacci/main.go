package main

import "fmt"

func main() {
	printFibonacciNumbers(8)
}

func printFibonacciNumbers(n int) {

	previous := 0
	next := 1
	sum := 0

	for i := 0; i <= n; i++ {
		fmt.Println("previous number is ", previous)
		sum = previous + next
		previous = next
		next = sum
	}
}
