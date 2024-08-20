package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(o *Point) float64 {
	return math.Sqrt((p.x-o.x)*(p.x-o.x) + (p.y-o.y)*(p.y-o.y))
}

func main() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(0, 3)
	distance := point1.Distance(point2)
	fmt.Println(distance)
}
