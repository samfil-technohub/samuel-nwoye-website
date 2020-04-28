package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", getEnvVar("PORT", "9000"), "Port to listen on")
	tmpl = template.New("")
)

func getEnvVar(desiredValue, defaultValue string) (value string) {
	value = os.Getenv(desiredValue)
	if value == "" {
		value = defaultValue
	}
	return
}

type Page struct {
	Title string
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	var err error
	_, err = tmpl.ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		log.Fatalf("Could not parse the template %v\n", err)
	}
	
	// initialize mux
	r := mux.NewRouter().StrictSlash(true)

	// define and server the static folder
	static := http.FileServer(http.Dir("templates"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", static))
	
	// define and server routes
	router := r.PathPrefix("/").Subrouter()
	router.HandleFunc("/", Index).Methods("GET")
	
	log.Printf("Server started and listening on port %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
