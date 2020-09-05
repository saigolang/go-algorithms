package main

import (
	"fmt"
	"reflect"
)

func main() {

	s1 := "abcde"

	s2 := "abfde"

	isStringOneAwayFromOtherStringSameLength(s1, s2)
}

func isStringOneAwayFromOtherStringSameLength(s1, s2 string) {

	if reflect.DeepEqual(s1, s2) {
		fmt.Println("result is false")
	}

	count := 0
	for i := 0; i <= len(s1)-1; i++ {

		if s1[i] != s2[i] {
			count++
			if count > 1 {
				fmt.Println(false)
				break
			}
		}
	}
	fmt.Println("They are one edit awy from each other")
}
