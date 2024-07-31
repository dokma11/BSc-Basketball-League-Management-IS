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

func (handler *InterviewHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	interviews, err := handler.InterviewService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var interviewResponseDTOs []model.InterviewResponseDTO
	for _, interview := range *interviews {
		var interviewResponseDTO model.InterviewResponseDTO
		interview.FromModel(&interviewResponseDTO)
		interviewResponseDTOs = append(interviewResponseDTOs, interviewResponseDTO)
	}

	json.NewEncoder(w).Encode(interviewResponseDTOs)
}

func (handler *InterviewHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var interviewResponseDTO model.InterviewResponseDTO
	interview.FromModel(&interviewResponseDTO)

	json.NewEncoder(w).Encode(interviewResponseDTO)
}

func (handler *InterviewHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
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

	var interviewResponseDTOs []model.InterviewResponseDTO
	for _, interview := range *interviews {
		var interviewResponseDTO model.InterviewResponseDTO
		interview.FromModel(&interviewResponseDTO)
		interviewResponseDTOs = append(interviewResponseDTOs, interviewResponseDTO)
	}

	json.NewEncoder(w).Encode(interviewResponseDTOs)
}

func (handler *InterviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	var interviewDTO model.InterviewCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&interviewDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	interview := &model.Interview{}
	interview.FromDTO(&interviewDTO)

	err := handler.InterviewService.Create(interview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
