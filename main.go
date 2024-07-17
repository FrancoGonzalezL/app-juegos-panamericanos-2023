package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const baseTempFile string = "templates/base.html"

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println("GET request received")
		tmpl := template.Must(template.ParseFiles(baseTempFile, "templates/home.html"))
		tmpl.Execute(w, nil)
	}

}

func main() {
	fmt.Println("Running server on port 8080")
	fmt.Print("http://127.0.0.1:8080\n\n")

	http.HandleFunc("/", home)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
