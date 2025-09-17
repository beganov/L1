package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y)))
}

func main() {
	a := NewPoint(-10.1, 9.1)
	b := NewPoint(-7, 5.1)
	fmt.Printf("Расстояние между точками: %.2f\n", a.Distance(b))
}
