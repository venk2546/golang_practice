package main

import "fmt"

type Employee struct {
	EmpName  string
	LastName string
	Salary   int
}

func main() {

	x1 := Employee{"ande", "Venkat", 30000}
	fmt.Println("Employee details", x1)

	x2 := Employee{
		EmpName:  "venkatesh",
		LastName: "ande",
		Salary:   30000,
	}

	fmt.Println("Emp detials:", x2)
	fmt.Println("only sal", x2.Salary)

}
