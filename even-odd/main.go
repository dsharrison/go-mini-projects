package main

import (
	"fmt"
)

func main() {

	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range ints {
		var numType string

		if num%2 == 0 {
			numType = "even"
		} else {
			numType = "odd"
		}

		fmt.Printf("%v is %v\n", num, numType)
	}
}
