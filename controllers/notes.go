package controllers

import (
	"basic-crud-api/models"
	"basic-crud-api/types"
	"basic-crud-api/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NoteResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note types.Note

	if r.Body == http.NoBody {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "missing request body",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	uuid, done := utils.ExtractUserFromJWT(w, r)
	if done {
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&note); err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	note.UserUUID = uuid
	if err := note.Validate(); err != nil {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	db := utils.Connection()
	payload := models.Note{
		Title:    note.Title,
		Body:     note.Body,
		UserUUID: note.UserUUID,
	}
	if err := db.Create(&payload).Error; err != nil {
		log.Print(err)
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "something went wrong",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	utils.CreateResponse(w, utils.ResponseBody{
		Message: "note created successfully",
		Code:    http.StatusCreated,
		Data:    NoteResponse{ID: payload.UUID, Title: payload.Title, Body: payload.Body, CreatedAt: payload.CreatedAt.Format("2006-01-02 15:04:05")},
	})
	return
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	uuid, done := utils.ExtractUserFromJWT(w, r)
	if done {
		return
	}

	db := utils.Connection()
	var notes []models.Note
	if err := db.Where("user_uuid = ?", uuid).Find(&notes).Error; err != nil {
		log.Print(err)
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "something went wrong",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	var data []NoteResponse
	for _, note := range notes {
		data = append(data, NoteResponse{ID: note.UUID, Title: note.Title, Body: note.Body, CreatedAt: note.CreatedAt.Format("2006-01-02 15:04:05")})
	}

	utils.CreateResponse(w, utils.ResponseBody{
		Message: "notes fetched successfully",
		Code:    http.StatusOK,
		Data:    data,
	})
	return
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	uuid, done := utils.ExtractUserFromJWT(w, r)
	if done {
		return
	}

	noteUUID := mux.Vars(r)["id"]
	if noteUUID == "" {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "missing note_uuid",
			Code:    http.StatusBadRequest,
			Data:    nil,
		})
		return
	}

	db := utils.Connection()

	count := db.Where("uuid = ? AND user_uuid = ?", noteUUID, uuid).Find(&models.Note{}).RowsAffected
	if count == 0 {
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "note not found",
			Code:    http.StatusNotFound,
			Data:    nil,
		})
		return
	}

	var note models.Note
	if err := db.Where("user_uuid = ? AND uuid = ?", uuid, noteUUID).First(&note).Error; err != nil {
		log.Print(err)
		utils.CreateResponse(w, utils.ResponseBody{
			Message: "something went wrong",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	utils.CreateResponse(w, utils.ResponseBody{
		Message: "note fetched successfully",
		Code:    http.StatusOK,
		Data:    NoteResponse{ID: note.UUID, Title: note.Title, Body: note.Body, CreatedAt: note.CreatedAt.Format("2006-01-02 15:04:05")},
	})
}
