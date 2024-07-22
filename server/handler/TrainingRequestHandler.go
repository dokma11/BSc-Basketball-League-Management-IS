package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TrainingRequestHandler struct {
	TrainingRequestService *service.TrainingRequestService
}

func NewTrainingRequestHandler(TrainingRequestService *service.TrainingRequestService) *TrainingRequestHandler {
	return &TrainingRequestHandler{TrainingRequestService: TrainingRequestService}
}

func (handler *TrainingRequestHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	trainingRequests, err := handler.TrainingRequestService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trainingRequests) // Proveriti samo da li valja
}

func (handler *TrainingRequestHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trainingRequest, err := handler.TrainingRequestService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if trainingRequest == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(trainingRequest)
}

func (handler *TrainingRequestHandler) GetAllBySenderID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trainingRequests, err := handler.TrainingRequestService.GetAllBySenderID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trainingRequests) // Proveriti samo da li valja
}

func (handler *TrainingRequestHandler) GetAllByReceiverID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trainingRequests, err := handler.TrainingRequestService.GetAllByReceiverID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trainingRequests) // Proveriti samo da li valja
}

func (handler *TrainingRequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var trainingRequest model.TrainingRequest
	if err := json.NewDecoder(r.Body).Decode(&trainingRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TrainingRequestService.Create(&trainingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TrainingRequestHandler) Update(w http.ResponseWriter, r *http.Request) {
	var trainingRequest model.TrainingRequest
	if err := json.NewDecoder(r.Body).Decode(&trainingRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TrainingRequestService.Update(&trainingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
