package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not Get method", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse failed %v", err)
	}
	if r.URL.Path != "/form" {
		http.Error(w, "not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Values: ")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %v\n", name)
	fmt.Fprintf(w, "address: %v\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //create a file server for index.html
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("starting the server at :8000\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { //start the server
		log.Fatal(err)
	}
}
