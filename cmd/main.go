package main

import (
	"aktest/internal/AK"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello AK!")
	})
	http.HandleFunc("/duck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gevorg!")
	})
	http.HandleFunc("/AK", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, AK.SymFunc())
	})
	http.ListenAndServe(":8080", nil)
}
