package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	add := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	add.Add(a, b)
	sub.Sub(a, b)
	mul.Mul(a, b)
	div.Div(a, b)
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", add.String())
	fmt.Printf("a - b = %s\n", sub.String())
	fmt.Printf("a * b = %s\n", mul.String())
	fmt.Printf("a / b = %s\n", div.String())
}
