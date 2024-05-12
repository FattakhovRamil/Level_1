package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	X float64
	Y float64
}

func (p *Point) Distance(other Point) float64 {

	dx := p.X - other.X
	dy := p.Y - other.Y

	return math.Sqrt(dx*dx + dy*dy)
}

func NewPoint(x float64, y float64) Point {
	return Point{X: x, Y: y}
}


func main() {
	first := NewPoint(2, 2)
	second := NewPoint(5, 5)
	result := first.Distance(second)
	fmt.Println(result)

}
