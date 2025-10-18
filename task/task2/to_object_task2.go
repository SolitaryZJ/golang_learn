package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}
func main() {
	e := Employee{Person: Person{Name: "zhang", Age: 18}, EmployeeID: 1}
	e.PrintInfo()
}
