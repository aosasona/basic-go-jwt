package controllers

import (
	"basic-crud-api/models"
	"basic-crud-api/types"
	"basic-crud-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body types.User
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&body)
	if err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
	}

	_, err = body.Validate()

	if err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	hashedPassword, err := body.HashPassword()

	if err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	payload := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  hashedPassword,
	}

	db := utils.Connection()

	res := db.Create(&payload)
	err = res.Error
	if err != nil {
		fmt.Print(err.Error())
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "something went wrong",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	msg := fmt.Sprintf("User %s created successfully", payload.FirstName)

	utils.CreateResponse(w, utils.ResponseBody{
		Message: msg,
		Code:    http.StatusCreated,
		Data: Response{
			ID:        payload.ID,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
		},
	})
	return
}
