package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// with the help of len function we can find the length of array
	// look we see the default value of string is empty string "" , there is no 2nd index value assigned yet
	// but the length of array is still 4
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list is: ", len(vegList))
}
