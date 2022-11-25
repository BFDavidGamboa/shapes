package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() (float64, error)
}

func main() {
	myTriangle := triangle{
		base:   10.0,
		height: 2.0,
	}
	mySquare := square{
		sideLength: 1.0,
	}
	printArea(myTriangle)
	printArea(mySquare)
}

func (t triangle) getArea() (float64, error) {
	const div float64 = 0.5
	if t.base <= 0 || t.height <= 0 {
		return 0, fmt.Errorf("no exists negative or zero areas for a triangle")
	}

	return (div * t.base * t.height), nil
}

func (s square) getArea() (float64, error) {
	if s.sideLength <= 0 {
		return 0, fmt.Errorf("no exists negative or zero areas for a square")
	}
	return (s.sideLength * s.sideLength), nil
}

func printArea(s shape) {

	area, err := s.getArea()
	if err != nil {
		fmt.Println(err)
	} else {
		out := fmt.Sprintf("%.2f\n", area)
		fmt.Printf("The shape has an area of %s\n", out)
	}

}
