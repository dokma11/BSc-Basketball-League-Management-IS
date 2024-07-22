package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type InterviewRequestHandler struct {
	InterviewRequestService *service.InterviewRequestService
}

func NewInterviewRequestHandler(InterviewRequestService *service.InterviewRequestService) *InterviewRequestHandler {
	return &InterviewRequestHandler{InterviewRequestService: InterviewRequestService}
}

func (handler *InterviewRequestHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	interviewRequests, err := handler.InterviewRequestService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interviewRequests) // Proveriti samo da li valja
}

func (handler *InterviewRequestHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	interviewRequest, err := handler.InterviewRequestService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if interviewRequest == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(interviewRequest)
}

func (handler *InterviewRequestHandler) GetAllBySenderID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	interviewRequests, err := handler.InterviewRequestService.GetAllBySenderID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interviewRequests) // Proveriti samo da li valja
}

func (handler *InterviewRequestHandler) GetAllByReceiverID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	interviewRequests, err := handler.InterviewRequestService.GetAllByReceiverID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(interviewRequests) // Proveriti samo da li valja
}

func (handler *InterviewRequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var interviewRequest model.InterviewRequest
	if err := json.NewDecoder(r.Body).Decode(&interviewRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.InterviewRequestService.Create(&interviewRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *InterviewRequestHandler) Update(w http.ResponseWriter, r *http.Request) {
	var interviewRequest model.InterviewRequest
	if err := json.NewDecoder(r.Body).Decode(&interviewRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.InterviewRequestService.Update(&interviewRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
