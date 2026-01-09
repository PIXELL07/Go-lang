package main

import (
	"fmt"
)

func main() {

	fmt.Println("Structs in Go")

	p := Person{"Alice", 30, "arun@gmail.com", true}
	fmt.Println(p)
	fmt.Printf("Person details are: %+v\n", p)
	p.GetStatus()
	p.NewEmail()
}

type Person struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u Person) GetStatus() {
	fmt.Println("User status is:", u.Status)

}

func (u Person) NewEmail() {
	u.Email = "test@dev"
	fmt.Println("User email is:", u.Email)
}
