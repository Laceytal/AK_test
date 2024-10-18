package main

import (
	"aktest/internal/handles"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handles.HandlOne)
	mux.HandleFunc("/duck", handles.HandlTwo)
	mux.HandleFunc("/AK", handles.HandlThree)
	http.ListenAndServe(":8080", mux)
}
