package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) (int, bool) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid, true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0, false
}

func main() {
	arr := []int{9, 2, 3, 4, 6, 2, 56, 5}
	sort.Ints(arr)
	fmt.Println(arr)
	var x = 5

	index, err := binarySearch(arr, x)
	if err == true {
		fmt.Printf("index = %d val = %d\n", index, x)
	} else {
		fmt.Println("not found")
	}
}
