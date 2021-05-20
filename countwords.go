package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Needles and pins Needles and pins Sew me a sail. To catch me the wind"
	words := strings.Fields(text)
	map_ := make(map[string]int)
	for _, word := range words {
		_, ok := map_[word]
		if !ok {
			map_[word] = 1
		} else {
			map_[word] = map_[word] + 1
		}
	}
	fmt.Println(map_)
}
