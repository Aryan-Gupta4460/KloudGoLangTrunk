package main

import (
	"fmt"
)

func Say_hello(msg string) {
	fmt.Println(msg)
}
func main() {
	Say_hello("Say Hello to All")
	func(name string) {
		fmt.Println(name)
	}("MY name iss Harry")

}
