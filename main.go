package main

import (
	"basic-crud-api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Index).Methods("GET")
	http.Handle("/", r)
}
