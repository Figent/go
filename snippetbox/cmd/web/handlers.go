package main

import(

	"fmt"
	"html/template"
	"net/http"
	"strconv"

)

func (app *application)home(w http.ResponseWriter, r*http.Request)  {

	if r.URL.Path != "/"{

		app.notFound(w)
		return

	}

	files := []string{

			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",

	}

	//use the template.ParseFiles()function to read template files into a template set
	//if there is a error we log the detailed error mesasage usinf the http.Error()function.

	ts, err := template.ParseFiles(files...)
	if err != nil {

		app.serverError(w, err)
		return

	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.

	err = ts.Execute(w, nil)
	if err != nil {

	app.serverError(w, err)
		
	}
	
	
}

func (app *application)showsnippet(w http.ResponseWriter, r*http.Request){

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	
	if err != nil || id < 1 {

		app.notFound(w)
		return

	}
	
	fmt.Fprintf(w, "Display snippet with the ID %d", id)

}

func (app *application)createsnippet(w http.ResponseWriter, r*http.Request){

	if r.Method != "POST" {

		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return

	}
	
	w.Write([]byte("Create snippet Please"))

}