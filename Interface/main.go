package main

import "fmt"

type shape interface {
	getArea() float64
}

type Triangle struct {
	height float64
	base   float64
}
type Square struct {
	sideLength float64
}

func main() {
	t := Triangle{1.2, 2.8}
	s := Square{2.6}
	printArea(t)
	printArea(s)

}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
