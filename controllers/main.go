package controllers

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Hello, world!"})
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	return
}
