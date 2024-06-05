package main

import (
	"fmt"
	"os"
)

func main() {
	var s string

	for i := len(os.Args) - 1; i > 0; i-- {
		s += os.Args[i] + " "
	}
	fmt.Print(s)
}
