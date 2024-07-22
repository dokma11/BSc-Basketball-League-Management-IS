package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type InterviewHandler struct {
	InterviewService *service.InterviewService
}

func NewInterviewHandler(InterviewService *service.InterviewService) *InterviewHandler {
	return &InterviewHandler{InterviewService: InterviewService}
}

func (handler *InterviewHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	interviews, err := handler.InterviewService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interviews) // Proveriti samo da li valja
}

func (handler *InterviewHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	interview, err := handler.InterviewService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if interview == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(interview)
}

func (handler *InterviewHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	interviews, err := handler.InterviewService.GetAllByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interviews) // Proveriti samo da li valja
}

func (handler *InterviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	var interview model.Interview
	if err := json.NewDecoder(r.Body).Decode(&interview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.InterviewService.Create(&interview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
