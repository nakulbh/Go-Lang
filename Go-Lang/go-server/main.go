package main

import (
	"fmt"
	"log"
	"net/http"
)

func formsHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %v\n", err)
		return
	}
	fmt.Fprintf(w, "Post request successfully \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)

	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formsHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting HTTP server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}

}
