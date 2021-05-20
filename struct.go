package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}
type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx, dy int) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func (s *Square) Move(dx, dy int) {
	s.Center.X = s.Center.X + dx
	s.Center.Y = s.Center.Y + dy
}
func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x, y, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("Invalid Length")
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}
func main() {

	square, _ := NewSquare(4, 5, 10)
	fmt.Println(square.Center)
	square.Move(2, 3)
	fmt.Println(square.Center)
	fmt.Println(square.Area())
}
