package main

import (
	"basic-jwt-api/controllers"
	"basic-jwt-api/models"
	"basic-jwt-api/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			panic("Unable to load .env file")
		}
	}
	port := ":" + os.Getenv("PORT")

	db := utils.Connection()
	err := db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		panic("Unable to migrate database")
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Headers("Accept", "application/json")

	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/auth/signup", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/auth/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/me", controllers.GetCurrentUser).Methods("GET")
	r.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	r.HandleFunc("/notes", controllers.GetAllNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", controllers.GetNote).Methods("GET")

	http.Handle("/", r)
	log.Printf("Server listening on port %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
