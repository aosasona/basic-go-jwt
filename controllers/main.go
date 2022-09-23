package controllers

import (
	"basic-crud-api/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.CreateResponse(w, utils.ResponseBody{
		Message: "Hey there!",
		Code:    http.StatusOK,
		Data:    nil,
	})
	return
}
