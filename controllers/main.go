package controllers

import (
	"basic-jwt-api/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.CreateResponse(w, utils.ResponseBody{
		Message: "Hey there!",
		Code:    http.StatusOK,
		Data:    nil,
	})
	return
}
