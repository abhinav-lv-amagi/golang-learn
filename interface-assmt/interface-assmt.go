package interfaceassmt

import "fmt"

func InterfaceAssmtEntry() {
	t := Triangle{
		height: 2,
		base:   4,
	}
	s := Square{
		sideLength: 4,
	}
	printArea(t)
	printArea(s)
}

type Shape interface {
	getArea() float64
}

type Triangle struct {
	height float64
	base   float64
}
type Square struct {
	sideLength float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println("Area of shape is:", s.getArea())
}
