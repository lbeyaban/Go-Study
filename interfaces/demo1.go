package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func school(s shape) {

	fmt.Println("Seklin alani : ", s.area())

}

func Demo1() {

	r := rectangle{width: 5, height: 4}
	school(r)

	t := triangle{height: 3, base: 2}
	school(t)

}
