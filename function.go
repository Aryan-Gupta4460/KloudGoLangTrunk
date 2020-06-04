package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}
func add3(a, b, c int) int {
	return a + b + c
}

func main() {
	ans := add(1, 3)
	fmt.Println(ans)
	ans33 := add3(2, 5, 7)
	fmt.Println(ans33)
}
