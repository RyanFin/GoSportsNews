package main

import (
	"RyanFin/GoSportsNews/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.NewsHandler)

	http.Handle("/", r)
	fmt.Println("GoSportsNews App - listening for requests on port :8080 ...")
	port := ":8080"
	http.ListenAndServe(port, r)
}
