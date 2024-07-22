package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RecruitHandler struct {
	RecruitService *service.RecruitService
}

func NewRecruitHandler(RecruitService *service.RecruitService) *RecruitHandler {
	return &RecruitHandler{RecruitService: RecruitService}
}

func (handler *RecruitHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	recruits, err := handler.RecruitService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recruits) // Proveriti samo da li valja
}

func (handler *RecruitHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	recruit, err := handler.RecruitService.GetByID(id)
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

func (handler *RecruitHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var recruit model.Recruit
	err := json.NewDecoder(req.Body).Decode(&recruit)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RecruitService.Create(&recruit)
	if err != nil {
		println("Error while creating a new recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *RecruitHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var recruit model.Recruit
	err := json.NewDecoder(req.Body).Decode(&recruit)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RecruitService.Update(&recruit)
	if err != nil {
		println("Error while updating recruit")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
