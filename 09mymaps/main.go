package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, My Maps!")

	languages := make(map[string]string) // make() allocates zeroed array & returns slice that refers to that array

	languages["en"] = "Hello"
	languages["es"] = "Hola"
	languages["fr"] = "Bonjour"

	fmt.Println("languages map:", languages)
	fmt.Println("es short form:", languages["es"])

	delete(languages, "fr")
	fmt.Println("languages map after deletion:", languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
