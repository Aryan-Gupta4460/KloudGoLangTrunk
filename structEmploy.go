package main

import "fmt"

type Employee struct {
	Employee_Id     int
	Employee_Name   string
	Employee_Design string
	Employee_Number []int
	Employee_Age    int
	Employee_Salary int
}

func main() {
	em1 := Employee{
		Employee_Id:     1,
		Employee_Name:   "Ram",
		Employee_Design: "Tester",
		Employee_Number: []int{
			8976523428, 6754348751},
		Employee_Age:    24,
		Employee_Salary: 20000,
	}
	em2 := Employee{
		Employee_Id:     2,
		Employee_Name:   "Prem",
		Employee_Design: "Devloper",
		Employee_Number: []int{9876523428},
		Employee_Age:    24,
		Employee_Salary: 25000,
	}
	em3 := Employee{
		Employee_Id:     3,
		Employee_Name:   "Ranjan",
		Employee_Design: "Analyst",
		Employee_Number: []int{7543523428, 4590763521},
		Employee_Age:    24,
		Employee_Salary: 20000,
	}
	em4 := Employee{
		Employee_Id:     4,
		Employee_Name:   "RajLaxmi",
		Employee_Design: "Senior Devloper",
		Employee_Number: []int{5632523428},
		Employee_Age:    24,
		Employee_Salary: 50000,
	}
	em5 := Employee{
		Employee_Id:     5,
		Employee_Name:   "Mrinal",
		Employee_Design: "Team Leader",
		Employee_Number: []int{7543542198},
		Employee_Age:    24,
		Employee_Salary: 60000,
	}

	fmt.Println(em1)
	fmt.Println(em2)
	fmt.Println(em3)
	fmt.Println(em4)
	fmt.Println(em5)
}
