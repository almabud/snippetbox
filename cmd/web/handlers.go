package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to 2 files.
	// It is important that file containing our base template must be
	// the first file in the slice.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// Use the template.parseFiles() function to read the template file
	// into a template set. If there's an error, we log the detailed error
	// message and use the http.Error() function to send a generic 500 Internal
	// server error response to the user.
	// We pass the slice of files using variadic parameter.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// We then use the Excute() method on the template set to write
	// the template content as the response body. The last parameter to
	// Execute() represents any dynamic data that we want to pass in, which
	// for now we'll leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Snippet %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Snippet created"))
}
