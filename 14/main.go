package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 42
	fmt.Println("Type", reflect.TypeOf(x))

	x = "hi"
	fmt.Println("Type", reflect.TypeOf(x))

	x = true
	fmt.Println("Type", reflect.TypeOf(x))

	x = make(chan int)
	fmt.Println("Type", reflect.TypeOf(x))

}
