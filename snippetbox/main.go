package main

import(

	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
		}

	w.Write([]byte("Hello from snippet"))

}

func showSnippet(w http.ResponseWriter, r *http.Request){

	//extract the value if the id from thr query string and try to convert it into an intiger using the srtconv,Atoi()function.

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	//if it cant be converted to an intiger or the value is less than one return 404 not found

	if err != nil || id < 1 {

		http.NotFound(w, r)

		return

	}

	//use the fmt.Fprintf() function to iterpolate our id with  with our response and and write it to the http.Responsewriter


	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

func createSnippet(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST" {

		w.Header().Set("Allow", "POST")

		http.Error(w, "Method Not Allowed", 405)

		return

	}

	w.Write([]byte("create a new snippet"))

}

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("starting server at port: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}