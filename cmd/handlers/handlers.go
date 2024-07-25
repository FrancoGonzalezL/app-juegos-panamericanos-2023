package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const baseTempFile string = "templates/base.html"

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println("GET request received for home")
		tmpl := template.Must(template.ParseFiles(
			baseTempFile,
			"templates/home.html",
		))
		tmpl.Execute(w, nil)
	}

}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET request received for stats")
		tmpl := template.Must(template.ParseFiles(
			baseTempFile, "templates/stats.html",
		))
		tmpl.Execute(w, nil)
	}

}
