package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the User Input Module!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok || err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input)
	fmt.Printf("Type of this rating is %T", input)

}
