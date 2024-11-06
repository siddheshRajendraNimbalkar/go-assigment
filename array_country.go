package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Enter the number of countries: ")
	fmt.Scan(&n)

	// Create a slice to store the country names
	countries := make([]string, n)

	// Accept country names from the user
	fmt.Println("Enter the country names:")
	for i := 0; i < n; i++ {
		fmt.Printf("Country %d: ", i+1)
		fmt.Scan(&countries[i])
	}

	// Sort the countries in descending order
	// sort.Sort(sort.Reverse(sort.StringSlice(countries)))
	sort.Strings(countries)

	// Display the sorted list of countries
	fmt.Println("\nCountries in descending order:")
	for _, country := range countries {
		fmt.Println(country)
	}
}
