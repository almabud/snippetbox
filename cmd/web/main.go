package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new command line flag with the name addr, a default value of
	// :4000 and some short help text explaining what the flag controls.
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assign it to the addr
	// variable. You need to call this before you use the addr variable
	// otherwise ti will always contain the default value. If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error message in the same way, but use
	// stderr as the destination and use the log.Lshortfile flag to include
	// relevant file name and line number.
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of our application struct, containing the
	// dependencis
	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	mux := app.routes()
	// Initialize a new http.Serveer struct. We set the addr and handler fields
	// so that the server uses the same network address and routes as before, and ste
	// the ErrorLog field so that the server now uses the custom errorLog loggeer in
	// the event of any problem
	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	// Write  the logg message using the new logger system.
	infoLog.Printf("Starting server on %s", *addr)
	//err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	//log.Println("Starting server on ", *addr)
	//err := http.ListenAndServe(*addr, mux)
	//log.Fatal(err)
	// To save the log into a file you need to run
	// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
}
