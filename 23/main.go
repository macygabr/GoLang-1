package main

import "fmt"

func main() {
	index := 2
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = remove(slice, index)
	fmt.Println(slice)
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
