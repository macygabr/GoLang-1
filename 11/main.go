package main

import (
	"fmt"
)

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 1}

	res := intersection(set1, set2)
	fmt.Println(res)
}

func intersection(set1, set2 []int) []int {
	result := make([]int, 0)

	for element := range set1 {
		for otherElement := range set2 {
			if element == otherElement {
				result = append(result, element)
				break
			}
		}
	}

	return result
}
