package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// define a handler function that writes a byte slice to the response
// writer. This will be registered as the handler for the "/" URL pattern
// r = pointer to a struct that holds information about the request

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{"./ui/html/partials/nav.tmpl.html", "./ui/html/base.tmpl.html", "./ui/html/home-page.tmpl.html"}
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// only possible to call it once on a response writer
		w.Header().Set("Allow", http.MethodPost)
		// make sure the header is set before writing the status code
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// this pattern of passing the writer is very common in Go
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
