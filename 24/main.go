package main

import (
	"fmt"
	"math"
)

/*
24. Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

// конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// нахождение расстояния
func (p *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

func main() {
	p1 := NewPoint(0, 1)
	p2 := NewPoint(10, 2)

	fmt.Println(p1.distance(p2))

}
