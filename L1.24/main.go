package main

import (
	"fmt"
	"math"
)

// L1.24
// Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64)
//  и конструктором. Расстояние рассчитывается по формуле между координатами двух точек.
// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

type Point struct { // Точки представлены в виде структуры Point
	x, y float64 //с инкапсулированными (приватными) полями x, y (типа float64)
}

func NewPoint(x, y float64) Point { // и конструктором
	return Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))) //Расстояние рассчитывается по формуле между координатами двух точек.
}

func main() {
	a := NewPoint(-10.1, 9.1)
	b := NewPoint(-7, 5.1)
	fmt.Printf("Расстояние между точками: %.2f\n", a.Distance(b))
}
