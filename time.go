package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println("Current date is: ", t.Format("02/01/2006"))
	fmt.Println("Current date is: ", t.Format("02-01-2006"))
	fmt.Println("Current date is: ", t.Format("Wednesday January 02 2006"))
}
