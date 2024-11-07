package main

import (
	"fmt"
	"sort"
)

type Student struct {
	rollno     int
	name       string
	percentage float64
}

func main() {
	var n int
	fmt.Println("Enter the number of students")
	fmt.Scan(&n)

	var s []Student
	s = make([]Student, n)

	for i := 0; i < n; i++ {
		var r int
		var name string
		var per float64

		fmt.Println("Enter rollno, name and percentage")
		fmt.Scan(&r, &name, &per)

		s[i].rollno = r
		s[i].name = name
		s[i].percentage = per

	}
	fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i].percentage > s[j].percentage
	})

	fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i].name < s[j].name
	})

	fmt.Println(s)
}
