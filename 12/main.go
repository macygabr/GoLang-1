package main

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]int)

	for _, element := range sequence {
		set[element] = set[element] + 1
	}

	fmt.Println(set)
}
