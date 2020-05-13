package controllers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"	
)


func ServeAPI() {
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