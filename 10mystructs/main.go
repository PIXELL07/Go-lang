package main

import (
	"fmt"
)

func main() {

	fmt.Println("Structs in Go")

	type Person struct {
		Name   string
		Age    int
		Email  string
		Status bool
	}

	p := Person{"Alice", 30, "arun@gmail.com", true}
	fmt.Println(p)
	fmt.Printf("Person details are: %+v\n", p)

}
