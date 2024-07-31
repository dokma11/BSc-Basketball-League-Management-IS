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

func (handler *TrainingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	trainings, err := handler.TrainingService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var trainingResponseDTOs []model.TrainingResponseDTO
	for _, training := range *trainings {
		var trainingResponseDTO model.TrainingResponseDTO
		training.FromModel(&trainingResponseDTO)
		trainingResponseDTOs = append(trainingResponseDTOs, trainingResponseDTO)
	}

	json.NewEncoder(w).Encode(trainingResponseDTOs)
}

func (handler *TrainingHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var trainingResponseDTO model.TrainingResponseDTO
	training.FromModel(&trainingResponseDTO)
	json.NewEncoder(w).Encode(trainingResponseDTO)
}

func (handler *TrainingHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
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

	var trainingResponseDTOs []model.TrainingResponseDTO
	for _, training := range *trainings {
		var trainingResponseDTO model.TrainingResponseDTO
		training.FromModel(&trainingResponseDTO)
		trainingResponseDTOs = append(trainingResponseDTOs, trainingResponseDTO)
	}

	json.NewEncoder(w).Encode(trainingResponseDTOs)
}

func (handler *TrainingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var trainingDTO model.TrainingCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&trainingDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	training := &model.Training{}
	training.FromDTO(&trainingDTO)

	err := handler.TrainingService.Create(training)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
