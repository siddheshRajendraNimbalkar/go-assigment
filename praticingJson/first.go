package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	RollNo int    `json:"rollno"`
}

func main() {
	var person Person
	person.Name = "SaiRam"
	person.Age = 20
	person.RollNo = 10

	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(jsonData))
	// reading json data
	var person1 Person

	err = json.Unmarshal(jsonData, &person1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person1)
}
