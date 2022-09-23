package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Message string
	Code    int
	Data    interface{}
}

func CreateResponse(w http.ResponseWriter, body ResponseBody) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(body.Code)
	code, message, data := body.Code, body.Message, body.Data
	var statusCode int
	if code != 0 {
		statusCode = code
	} else {
		statusCode = http.StatusOK
	}
	err := json.NewEncoder(w).Encode(map[string]any{
		"message": message,
		"code":    statusCode,
		"data":    data,
	})
	if err != nil {
		err := json.NewEncoder(w).Encode(map[string]any{
			"message": message,
			"code":    statusCode,
			"data":    data,
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
	return
}
