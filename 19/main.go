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
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
