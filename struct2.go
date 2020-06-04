package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {
	a := Bird{}
	a.Name = "Owl"
	a.Origin = "India"
	a.Speed = 80.4
	a.CanFly = true

	fmt.Println(a)
}
