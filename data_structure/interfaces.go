package main

import "fmt"

type animal interface {
	speak() string
}
type cat struct {
	sound string
}

func (c cat) speak() string {
	return c.sound
}

type dog struct {
	sound string
}

func (d dog) speak() string {
	return d.sound
}
func main() {
	c := cat{"meaw"}
	d := dog{"bow boww"}
	animals := []animal{c, d}
	for _, animal := range animals {
		fmt.Println(animal.speak())
	}
}
