package main

import (
	"fmt"
)

func main() {
	StatePopulation := map[string]int{
		"india": 1369,
		"china": 3000,
		"nepal": 100,
		"japan": 5000,
	}
	if pop, ok := StatePopulation["india"]; ok {
		fmt.Println(pop)
	}
}
