package main

import (
	"fmt"
)

func retun_annomus_func() func(string) {
	return func(name string) {
		fmt.Println(name)
	}
}
func int_seQ() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func Say_hello(msg string) {
	fmt.Println(msg)
}
func main() {
	Say_hello("Say Hello to All")
	func(name string) {
		fmt.Println(name)
	}("MY name iss Harry")
	print_func := retun_annomus_func()
	print_func("Hey we are to close")

	next_int := int_seQ()
	fmt.Println(next_int())
	fmt.Println(next_int())

	new_int2 := int_seQ()
	fmt.Println(new_int2())
}
