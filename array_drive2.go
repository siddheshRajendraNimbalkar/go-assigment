package main

import "fmt"

func arryaDrive() {
	fmt.Println("--------------------------------------")
	fmt.Println("1. Sum of diagonal elements")
	fmt.Println("2. Sum of upper diagonal elements")
	fmt.Println("3. Sum of lower diagonal elements")
	fmt.Println("4. Exit")
	fmt.Println("--------------------------------------")
}

func main() {
	var arr [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for true {
		var cho int
		arryaDrive()
		fmt.Println("Enter your choice")
		fmt.Scan(&cho)
		switch cho {
		case 1:
			fmt.Println("Sum of diagonal elements")
			var sum int = 0
			for i := 0; i < len(arr); i++ {
				sum += arr[i][i]
			}
			fmt.Println("Sum is ", sum)
			break
		case 2:
			fmt.Println("Sum of upper diagonal elements")
			var sum int = 0
			for i := 0; i < len(arr); i++ {
				for j := 0; j < len(arr); j++ {
					if i < j {
						sum += arr[i][j]
					}
				}
			}
			fmt.Println("Sum is ", sum)
			break
		case 3:
			fmt.Println("Sum of lower diagonal elements")
			var sum int = 0
			for i := 0; i < len(arr); i++ {
				for j := 0; j < len(arr); j++ {
					if i > j {
						sum += arr[i][j]
					}
				}
			}
			fmt.Println("Sum is ", sum)
			break
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
			break
		}
	}
}
