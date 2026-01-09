package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gh_1234"

func main() {
	fmt.Println("welcome to 19 urls")
	fmt.Println(myurl)

	//parsing the url

	result, _ := url.Parse(myurl)

	//fmt.Println(result.Scheme)
	//fmt.Println(result.Host)
	//fmt.Println(result.Path)
	//fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	//looping through the map
	// for _, which defines not worried about the key or index we only want the value

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//constructing the url
	// passing on a reference to &url.URL

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "coursename=css&paymentid=gh_5678",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
