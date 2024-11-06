package main

import "fmt"

func main() {
	fmt.Println("Printing the area and perimeter of a rectangle")
	var length, breadth float64
	fmt.Println("Enter the length of the rectangle")
	fmt.Scan(&length)
	fmt.Println("Enter the breadth of the rectangle")
	fmt.Scan(&breadth)
	var area = length * breadth
	var perimeter = 2 * (length + breadth)
	fmt.Println("Area of the rectangle is", area)
	fmt.Println("Perimeter of the rectangle is", perimeter)
}
