package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frgonzal/app-juegos-panamericanos-2023/handlers"
)

func main() {
	fmt.Println("Running server on port 8080")
	fmt.Print("http://127.0.0.1:8080\n\n")

	http.HandleFunc("/stats/", handlers.StatsHandler)
	http.HandleFunc("/", handlers.HomeHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
