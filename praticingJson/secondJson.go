package main

import (
	Packagejson "ass/packageJson"
	"fmt"
)

type Name struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}

func main() {

	a, err := Packagejson.SendJson()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(a)
}
