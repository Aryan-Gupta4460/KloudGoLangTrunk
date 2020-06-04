package main

import (
	"fmt"
)

func multi(num ...int) int {
	total := 1
	for _, num := range num {

		total *= num
	}
	return total
}
func main() {
	fmt.Println(multi(1, 3, 6, 2))
	nums := []int{3, 6, 2, 1, 5}
	fmt.Println(multi(nums...))
}
