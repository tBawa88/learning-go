package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Let's make a web server")
	//with this it will look inside the static directory to serve the incoming http requests
	fileServer := http.FileServer(http.Dir("./static"))

	//first argument is the route, second is a handler type
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server", err)
	}

	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, r.Method+" requests not allowed to this route", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v \n", err)
		return
	}

	fmt.Fprintf(w, "POST request successfull")
	username := r.FormValue("username")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Username: %s , Email: %s", username, email)
}
