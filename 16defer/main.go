package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("from")
	defer fmt.Println("Go")
	fmt.Println("hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
}
