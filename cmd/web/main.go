package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static"
	// directory. This is relative path.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as
	// the handler for all URL paths that start with /static/.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/", &home{})
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Define a new command line flag with the name addr, a default value of
	// :4000 and some short help text explaining what the flag controls.
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assign it to the addr
	// variable. You need to call this before you use the addr variable
	// otherwise ti will always contain the default value. If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	log.Println("Starting server on ", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
