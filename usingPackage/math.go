package main

import (
	"ass/myPackage"
	"fmt"
)

func main() {
	var a, b int64 = 2, 3
	fmt.Println("a:", a, " b:", b)
	fmt.Println("add: ", myPackage.Add(a, b))
	fmt.Println("sub: ", myPackage.Sub(a, b))
	fmt.Println("mul: ", myPackage.Mul(a, b))
	fmt.Println("div: ", myPackage.Div(a, b))
}
