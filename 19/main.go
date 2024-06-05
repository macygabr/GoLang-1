package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[len(os.Args)-1]

	fmt.Printf("%s", reverse(s[:]))
}

func reverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
