package main

import "fmt"

type react struct {
	Width  float32
	height float32
}

func (r react) area() float32 {
	return r.Width * r.height
}

func (r *react) perimeter() float32 {
	return 2 * r.Width * 2 * r.height
}

func main() {
	r := react{Width: 10, height: 2}
	fmt.Println("Area", r.area())

	r_ptr := &r
	fmt.Println("Perimeter", r_ptr.perimeter())
}
