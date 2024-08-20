package main

import "fmt"

func main() {
	testString := "aabcd"
	fmt.Println(checkString(testString))
}

func checkString(testString string) bool {
	charMap := make(map[rune]bool)

	for _, char := range testString {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}
