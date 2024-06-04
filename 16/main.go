package main

import "fmt"

type inter interface {
	float64 | float32 | int32 | int16 | int64 | int8
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(arr)
	quickHoaraSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickHoaraSort[a inter](arr []a, start int, end int) {
	if start >= end {
		return
	}

	rightStart := partOfSortHoara(arr, start, end)
	quickHoaraSort(arr, start, rightStart-1)
	quickHoaraSort(arr, rightStart, end)
}

func partOfSortHoara[a inter](arr []a, left int, right int) int {
	pivot := arr[left+(right-left)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}
