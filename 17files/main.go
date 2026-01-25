package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to 17 files example")
	content := "this needs to go in a file"

	file, err := os.Create("./myfile.txt")

	// if err != nil {   // panic() , recover()
	//	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilError(err)

	fmt.Println("text data inside file is: ", string(databyte))
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
