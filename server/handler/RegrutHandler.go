package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RegrutHandler struct {
	RegrutService *service.RegrutService
}

func NewRegrutHandler(regrutService *service.RegrutService) *RegrutHandler {
	return &RegrutHandler{RegrutService: regrutService}
}

func (handler *RegrutHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	recruits, err := handler.RegrutService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recruits) // Proveriti samo da li valja
}

func (handler *RegrutHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	recruit, err := handler.RegrutService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recruit == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(recruit)
}

func (handler *RegrutHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var recruit model.Regrut
	err := json.NewDecoder(req.Body).Decode(&recruit)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegrutService.Create(&recruit)
	if err != nil {
		println("Error while creating a new recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *RegrutHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var recruit model.Regrut
	err := json.NewDecoder(req.Body).Decode(&recruit)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegrutService.Update(&recruit)
	if err != nil {
		println("Error while updating recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
