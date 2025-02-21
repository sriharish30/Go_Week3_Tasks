package main

import "fmt"

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func main() {
	c := circle{radius: 3}
	r := rectangle{length: 4, breadth: 5}
	shapes := []shape{c, r}

	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.area())
	}
}
