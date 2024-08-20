package main

import "fmt"

func main() {
	x := 5
	y := 10
	fmt.Println(x, y)

	x ^= y
	y ^= x
	x ^= y

	fmt.Println(x, y)
}
