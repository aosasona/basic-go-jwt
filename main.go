package main

import (
	"basic-crud-api/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load .env file")
	}
	port := ":" + os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Index).Methods("GET")
	http.Handle("/", r)
	log.Printf("Server listening on port %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
