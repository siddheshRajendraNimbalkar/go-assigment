package main

import "fmt"

func main() {
	var a []int
	var s int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&s)

	for i := 0; i < s; i++ {
		var x int
		fmt.Println("Enter the element")
		fmt.Scan(&x)
		a = append(a, x)
	}
	fmt.Println("The array in reverse order is")
	for i := s - 1; i >= 0; i-- {
		fmt.Print("\t", a[i])
	}
}
