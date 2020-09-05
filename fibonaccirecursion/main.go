package main

import "fmt"

func main() {
	fmt.Println("fibnacci recusrion result is", fibonacciRecursion(4))

}

func fibonacciRecursion(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}
