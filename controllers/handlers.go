package controllers

import (
	"flag"
	"html/template"
	"net/http"
	"os"

	_"github.com/samfil-technohub/samuel-nwoye-website/models"
)

var (
	port = flag.String("port", getEnvVar("PORT", "9000"), "Port to listen on")
	tmpl = template.New("")
)

// option to pass values as environment variable
func getEnvVar(desiredValue, defaultValue string) (value string) {
	value = os.Getenv(desiredValue)
	if value == "" {
		value = defaultValue
	}
	return
}

// index page function
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}