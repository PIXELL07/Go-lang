package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list is: ", len(vegList))
}
