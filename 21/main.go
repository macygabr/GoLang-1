package main

import (
	"fmt"
	"math"
)

type Square interface {
	Area() float64
}

type Circle interface {
	Circumference() float64
}

type MySquare struct {
	side float64
}

func (s *MySquare) Area() float64 {
	return s.side * s.side
}

type MyCircle struct {
	radius float64
}

func (c *MyCircle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

type SquareToCircleAdapter struct {
	square Square
}

func (a *SquareToCircleAdapter) Circumference() float64 {
	area := a.square.Area()
	radius := math.Sqrt(area / math.Pi)
	return 2 * math.Pi * radius
}

func main() {
	mySquare := &MySquare{side: 4}
	adapter := &SquareToCircleAdapter{square: mySquare}
	fmt.Printf("Circumference of the equivalent circle: %.2f\n", adapter.Circumference())
}
