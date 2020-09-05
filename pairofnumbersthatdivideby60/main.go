package main

import "fmt"

func main() {
	n := []int{30, 60, 120, 40, 80}

	pairOfNumbersThatDivideBy60(n)
}

func pairOfNumbersThatDivideBy60(n []int) {

	result := 0
	sample := make(map[int]int)
	for _, r := range n {
		if _, ok := sample[r]; !ok {
			sample[r] = r
			if r%60 == 0 {
				result++
			}
		} else {
			for k, _ := range sample {
				if sample[k]%60 == 0 {
					result++
				}
			}
		}
	}
	fmt.Println("result is ", result)
}
