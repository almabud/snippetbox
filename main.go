package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux,
	// then register the home function as the handler for "/" url pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenServe() function to start a new web server. We
	// pass in two parameters: the TCP network address to listen.
	// and the servemux we just created. If http.ListenAndServe() return an
	// error we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by the http.ListenAndServe() is always no-nil.
	log.Println("Listening on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
