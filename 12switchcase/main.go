package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value 1 and you can open")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("You can move 4 steps")
		fallthrough // by using fallthrough it stops the code and doesnt print the nxt line
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and roll the dice again")
	default:
		fmt.Println("What was that?")
	}
}
