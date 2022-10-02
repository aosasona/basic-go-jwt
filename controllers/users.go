package controllers

import (
	"basic-crud-api/models"
	"basic-crud-api/types"
	"basic-crud-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type SignUpResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
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

	if payload.CheckAlreadyExists(db) {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "user already exists",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	res := db.Create(&payload)
	if err = res.Error; err != nil {
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
		Data: SignUpResponse{
			ID:        payload.UUID,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
		},
	})
	return
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body types.User
	if r.Body == http.NoBody {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "all fields are required",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&body); err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	if _, err := body.ValidateLoginData(); err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	db := utils.Connection()

	var user models.User

	db.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "invalid credentials",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	if err := body.ComparePassword(user.Password); err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "invalid credentials",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	token, err := utils.GenerateJWT(user.UUID)

	if err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "something went wrong",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	res := LoginResponse{
		Email: user.Email,
		Token: token,
	}

	utils.CreateResponse(w, utils.ResponseBody{
		Message: "welcome back",
		Code:    http.StatusOK,
		Data:    res,
	})
	return
}
