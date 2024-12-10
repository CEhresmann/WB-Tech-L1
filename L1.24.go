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

func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distance := point1.Distance(point2)
	fmt.Printf("Расстояние между точками (%.0f, %.0f) и (%.0f, %.0f) равно %.2f\n",
		point1.x, point1.y, point2.x, point2.y, distance)
}
