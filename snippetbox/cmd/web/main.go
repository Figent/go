package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"alexedwards.net/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//application Struct

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	//define a new command line-flag and assing it the default value:4000 and an explanantion of what command line flag does

	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
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

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{

		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct.
	srv := &http.Server{

		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), //call the new app rolls

	}

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it.

	infoLog.Printf("starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
