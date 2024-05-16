package main

import (
	"net/http"
	"strconv"
)

// define a handler function that writes a byte slice to the response
// writer. This will be registered as the handler for the "/" URL pattern
// r = pointer to a struct that holds information about the request

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, World!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display the snippet"))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// only possible to call it once on a response writer
		w.Header().Set("Allow", http.MethodPost)
		// make sure the header is set before writing the status code
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// this pattern of passing the writer is very common in Go
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
