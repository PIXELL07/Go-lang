package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

// Modules and dependencies
// * go mod download
// * go mod list all
// * go mod edit
// * go mod graph
// * go mod init
// * go mod tidy
// * go mod vendor
// * go mod verify
// * go mod why
// * go version -m
// * go clean -modcache
// * go run  -mod=versions main.go

// * Module commands outside a module
// go work init
// go work edit
// go work use
// go work sync
