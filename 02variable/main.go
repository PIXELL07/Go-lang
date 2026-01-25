package main

import "fmt"

const loginToken string = "arun"

func main() {
	var username string = "pee"
	fmt.Println(username)
	fmt.Printf("variable username is of type: %T \n", username)

	var isLoggesIn bool = true
	fmt.Println(isLoggesIn)
	fmt.Printf("variable username is of type: %T \n", isLoggesIn)

	var aaiii uint8 = 255
	fmt.Println(aaiii)
	fmt.Printf("variable username is of type: %T \n", aaiii)

	var sandy float64 = 255.4557834717491208742
	fmt.Println(sandy)
	fmt.Printf("variable username is of type: %T \n", sandy)

	// default value and some alias
	var yooo int
	fmt.Println(yooo)
	fmt.Printf("variable username is of type: %T \n", yooo)

	// implicit type
	var hiiii = "sandyali"
	fmt.Println(hiiii)

	// no var style
	noOfUsers := "saaandaliyannn"
	fmt.Println(noOfUsers)

	fmt.Println("login token is: ", loginToken)
}
