package controllers

import (
	"basic-crud-api/types"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body types.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

}
