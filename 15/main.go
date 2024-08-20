package main

var justString *string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = &v
}

func main() {
	someFunc()
}
