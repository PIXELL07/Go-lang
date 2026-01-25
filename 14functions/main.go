package main

import "fmt"

func main() {
	fmt.Println("welcome to Functions!") // Functions are collection of statements
	greeter()

	result := add(3, 4)

	fmt.Println("Result is:", result)

	proResult, Message := proAdd(1, 2, 3, 4, 5)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Message is:", Message)
}

func add(a int, b int) int {
	return a + b
}

func proAdd(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}
	return total, "hi bro"
}
func greeter() {
	fmt.Println("Hello from greeter function")
}
