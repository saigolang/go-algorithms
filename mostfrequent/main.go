package main

import "fmt"

func main() {

	input := []int{1, 1, 2, 3, 4, 4, 5}

	mostFrequent(input)

}

func mostFrequent(n []int) {
	sample := make(map[int]int)

	maxItem := 0
	maxItemCount := 1
	for _, r := range n {
		if _, ok := sample[r]; !ok {
			sample[r] = 1
		} else {
			sample[r]++
		}
		if sample[r] > maxItemCount {
			maxItem = r
			maxItemCount++
		}
	}

	fmt.Println("max item is ", maxItem)
	fmt.Println("max item count is ", maxItemCount)
}
