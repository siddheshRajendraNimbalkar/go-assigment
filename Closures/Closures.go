package main

import "fmt"

func addSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	a := addSeq()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	b := addSeq()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
