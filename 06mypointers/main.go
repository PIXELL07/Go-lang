package main

import "fmt"

func main() {
	fmt.Println("Hello, my pointers!")

	myNumber := 42

	var ptr = &myNumber //address of myNumber

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of pointer is", *ptr) //dereferencing pointer, value of address

	*ptr = *ptr + 8
	fmt.Println("new value is ", myNumber)
}
