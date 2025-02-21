package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "dog: Woof!"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "cat: Meow!"
}

type Cow struct{}

func (c Cow) Sound() string {
	return "cow: Moo!"
}

func main() {

	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	animals := []Animal{dog, cat, cow}

	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}
