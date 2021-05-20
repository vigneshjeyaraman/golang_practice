package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}
type Circle struct {
	R float64
}

type Shapes interface {
	Area() float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}
func (c *Circle) Area() float64 {
	return 2.0 * math.Pi * c.R
}
func sumAreas(shapes []Shapes) float64 {
	sum := 0.0
	for _, shape := range shapes {
		sum = sum + shape.Area()
	}
	return sum
}

func main() {
	c := &Circle{3}
	s := &Square{5}
	shapes := []Shapes{c, s}
	area := sumAreas(shapes)
	fmt.Println(area)
}
