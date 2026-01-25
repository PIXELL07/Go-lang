package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, myslices!")

	var fruitList = []string{"apple", "banana", "grapes", "mango"} //slice of string
	fmt.Printf("type of fruit is %T\n", fruitList)

	fruitList = append(fruitList, "orange", "kiwi") // append: used for adding
	fmt.Println(fruitList)

	fruitList = fruitList[1:3] // last range is always non-inclusive , so it will print index 1 and 2
	fmt.Println(fruitList)

	// Appending to a nil slice
	var boss []string
	boss = append(boss, "hello", "world")
	fmt.Println("Appended to nil slice:", boss) // Output: Appended to nil slice: [hello world]

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 678
	highScores[3] = 890

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("highest score is:", highScores[len(highScores)-1])
	fmt.Println("lowest score is:", highScores[0])
	fmt.Println("length of highScores is:", len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slice based on index
	var courses = []string{"reactjs", "javascript", "html", "css", "angular"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // Print the updated courses slice

}
