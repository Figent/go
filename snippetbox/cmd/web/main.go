package main

import (

	"flag"
	"log"
	"net/http"
	"os"

)

//application Struct

type application struct {

	errorLog *log.Logger
	infoLog *log.Logger

}

func main(){

	//define a new command line-flag and assing it the default value:4000 and an explanantion of what command line flag does

	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.

	flag.Parse()

	//using the log.New()functoin we create a logger for writing informational messages with three parameters.
	//parameter one:destination to write the logs (os.Stdout)
	//parameter two:a string prefix for message (INFO followed by a tab)
	//parameter three:flags to indicate what additional information is required.

	infoLog := log.New(os.Stdout, "INFO	", log.Ldate|log.Ltime)

	//using the log.New()functoin we create a logger for writing Error informational messages with three parameters.
	//parameter one:destination to write the logs (os.Stderr)
	//parameter two:a string prefix for message (ERROR followed by a tab)
	//parameter three:flags to indicate what additional information is required.

	errorLog := log.New(os.Stderr, "ERROR	", log.Ldate|log.Ltime|log.Llongfile)
	app := &application{

		errorLog: errorLog,
		infoLog: infoLog,

	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showsnippet)
	mux.HandleFunc("/snippet/create", app.createsnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct.
	srv := &http.Server{

		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,

	}

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it.

	infoLog.Printf("starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}