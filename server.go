package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server at 8080")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error : ", err)
	}	
}