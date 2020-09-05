package main

import "fmt"

func main() {

	input := []string{"A", "B", "A", "B"}

	firstRecurringCharacter(input)
}

func firstRecurringCharacter(n []string) {

	testMap := make(map[string]int)
	for _, r := range n {

		_, ok := testMap[r]

		if !ok {
			testMap[r] = 1
		} else {
			testMap[r]++
		}
		if testMap[r] > 1 {
			fmt.Println("most occured value is ", r)
			break
		}
	}
}
