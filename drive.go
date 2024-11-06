package main

import "fmt"

func drive() {
	fmt.Println("-------------------------------------")
	fmt.Println("1. Calc the volume of cylinder")
	fmt.Println("2. Find the factorial of a number")
	fmt.Println("3. Check the no. is Armstrong or not")
	fmt.Println("4. Exit")
	fmt.Println("-------------------------------------")
}

func main() {
	fmt.Println("Drive Program")
	for true {
		drive()
		var cho int
		fmt.Println("Enter your choice")
		fmt.Scan(&cho)
		switch cho {
		case 1:
			fmt.Println("Calc the volume of cylinder")
			var r, h float64
			fmt.Println("Enter the radius ")
			fmt.Scan(&r)
			fmt.Println("Enter the height ")
			fmt.Scan(&h)
			fmt.Println("The volume of cylinder is ", 3.14*r*r*h)
			break
		case 2:
			fmt.Println("Find the factorial of a number")
			var n int
			sum := 1
			fmt.Println("Enter the number")
			fmt.Scan(&n)
			for i := 1; i <= n; i++ {
				sum *= i
			}
			fmt.Println("The factorial is ", sum)
			break
		case 3:
			fmt.Println("Check the no. is Armstrong or not")
			var n int
			fmt.Println("Enter the number")
			fmt.Scan(&n)
			var temp = n
			var sum = 0
			for n != 0 {
				digit := n % 10
				sum += digit * digit * digit
				n = n / 10
			}
			if sum == temp {
				fmt.Println("The no. is Armstrong")
			} else {
				fmt.Println("The no. is not Armstrong")
			}
			break

		case 4:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
