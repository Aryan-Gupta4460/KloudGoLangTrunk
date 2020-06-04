package main

import (
	"fmt"
)

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}
func main() {
	add_res, sub_res := add_sub(1, 2)
	fmt.Println("Add_Res", add_res, "Sub_Res", sub_res)
}
