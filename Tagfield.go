package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name   string `required max :"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Bird{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field, t)
}
