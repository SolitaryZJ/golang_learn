package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 5}
	shapes := []Shape{r, c}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
		fmt.Println(shape.Perimeter())
	}
}
