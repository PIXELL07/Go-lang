package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("welcome to web requests in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the body

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err) // handles, sometimes breaks n err
	}
	content := string(databytes)
	fmt.Println(content)
}
