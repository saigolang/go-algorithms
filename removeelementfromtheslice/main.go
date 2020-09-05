package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 5}

	// remove element from the slice with changing the order of slice
	index := 3
	s[index] = s[len(s)-1]

	fmt.Println("slice after removing the element is ", s[:len(s)-1])

	// remove element from the slice without changing the order of a slice
	index = 4
	copy(s[index:], s[index+1:])

	fmt.Println("slice after removing the element is ", s[:len(s)-1])

}
