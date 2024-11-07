package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var a []Person
	var s int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&s)

	for i := 0; i < s; i++ {
		var name string
		var age int
		var salary float64
		fmt.Println("Enter the name of the person")
		fmt.Scan(&name)
		fmt.Println("Enter the age of the person")
		fmt.Scan(&age)
		fmt.Println("Enter the salary of the person")
		fmt.Scan(&salary)
		a = append(a, Person{Name: name, Age: age, Salary: salary})
	}

	for i := 0; i < s; i++ {
		fmt.Println("Name: ", a[i].Name)
		fmt.Println("Age: ", a[i].Age)
		fmt.Println("Salary: ", a[i].Salary)
	}
}
