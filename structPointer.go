package main

import (
	"fmt"
)

type Employe struct {
	First_Name string
	Last_name  string
	Id         int
}

func main() {
	fmt.Println(Employe{First_Name: "Aryan", Last_name: "Gupta", Id: 1})
	fmt.Println(Employe{"Ram", "Roy", 2})
	emp := Employe{First_Name: "Abhishek", Last_name: "Dutta", Id: 3}

	fmt.Println(emp.First_Name)
	fmt.Println(emp.Last_name)
	fmt.Println(emp.Id)
	fmt.Println(emp)

	emp_ptr := &emp
	fmt.Println(emp_ptr.First_Name)
	emp_ptr.First_Name = "Alok"
	fmt.Println(emp_ptr.First_Name)
	fmt.Println(emp.First_Name)
}
