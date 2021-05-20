package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max_ := -1
	for _, value := range nums {
		if value > max_ {
			max_ = value
		}
	}
	fmt.Println("Max number is:- ", max_)
}
