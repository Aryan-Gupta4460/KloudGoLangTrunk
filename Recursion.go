package main

import "fmt"

func Recur(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Recur(n-1)
}
func main() {
	fmt.Println(Recur(4))
}
