package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frgonzal/juegos-panamericanos/cmd/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Running server on port 8080")
	fmt.Print("http://127.0.0.1:8080\n\n")

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/stats/", handlers.StatsHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
