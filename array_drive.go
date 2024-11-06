package main

import "fmt"

func arryaDrive() {
	fmt.Println("--------------------------------------")
	fmt.Println("1. Addition of Array")
	fmt.Println("2. Multiplication of Array")
	fmt.Println("3. Transpose of Array")
	fmt.Println("4. Exit")
	fmt.Println("--------------------------------------")
}

func main() {
	fmt.Println("Array Drive Program")
	arr := []int{1, 2, 3, 4, 5}
	arr1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	var arr2 [][]int

	for true {
		arryaDrive()
		var cho int
		fmt.Println("Enter your choice")
		fmt.Scan(&cho)
		switch cho {
		case 1:
			fmt.Println("Addition of Array")
			var sum int = 0
			for i := 0; i < len(arr); i++ {
				sum += arr[i]
			}
			fmt.Println("Sum is ", sum)
			break
		case 2:
			fmt.Println("Multiplication of Array")
			var mul int = 1
			for i := 0; i < len(arr); i++ {
				mul *= arr[i]
			}
			fmt.Println("Multiplication is ", mul)
			break
		case 3:
			fmt.Println("Transpose of Array")
			rows, coln := len(arr1), len(arr1[0])
			arr2 = make([][]int, coln)
			for i := 0; i < len(arr2); i++ {
				arr2[i] = make([]int, rows)
			}
			for i := 0; i < len(arr1[0]); i++ {
				for j := 0; j < len(arr1); j++ {
					arr2[i][j] = arr1[j][i]
				}
			}
			for i := 0; i < len(arr2); i++ {
				for j := 0; j < len(arr2[i]); j++ {
					fmt.Print(arr2[i][j], "\t")
				}
				fmt.Println()
			}
			break
		case 4:
			fmt.Println("Exit")
			return
		}
	}

}
