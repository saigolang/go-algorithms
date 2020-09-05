package main

import "fmt"

func main() {
	n := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}

	sumOfPairsEqualTo11(n)
}

func sumOfPairsEqualTo11(n []int) {

	l := 0
	r := len(n) - 1

	for l < r {

		if n[l]+n[r] > 11 {
			r--
		} else if n[l]+n[r] < 11 {
			l++
		} else {
			fmt.Println("pairs are ", n[l], n[r])
			l++
		}
	}

}
