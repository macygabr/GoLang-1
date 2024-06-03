package main

import (
	"fmt"
)

func main() {
	var mask int64
	var x int64
	var i int
	var byt int
	mask = 1
	x = 10000
	i = 5
	byt = 1

	fmt.Printf("%064b\n", x)
	mask = mask << i
	fmt.Printf("%064b\n", mask)
	if byt == 1 {
		x = x | mask
		fmt.Printf("|\n")
	} else {
		x = x &^ mask
		fmt.Printf("&^\n")
	}
	fmt.Printf("%064b\n", x)
}
