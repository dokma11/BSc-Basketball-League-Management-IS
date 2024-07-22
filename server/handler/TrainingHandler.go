package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TrainingHandler struct {
	TrainingService *service.TrainingService
}

func NewTrainingHandler(TrainingService *service.TrainingService) *TrainingHandler {
	return &TrainingHandler{TrainingService: TrainingService}
}

func (handler *TrainingHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	trainings, err := handler.TrainingService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trainings) // Proveriti samo da li valja
}

func (handler *TrainingHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	training, err := handler.TrainingService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if training == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(training)
}

func (handler *TrainingHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trainings, err := handler.TrainingService.GetAllByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trainings) // Proveriti samo da li valja
}

func (handler *TrainingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var training model.Training
	if err := json.NewDecoder(r.Body).Decode(&training); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TrainingService.Create(&training)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
