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

func (handler *InterviewRequestHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	interviewRequests, err := handler.InterviewRequestService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var interviewRequestResponseDTOs []model.InterviewRequestResponseDTO
	for _, interviewRequest := range *interviewRequests {
		var interviewRequestResponseDTO model.InterviewRequestResponseDTO
		interviewRequest.FromModel(&interviewRequestResponseDTO)
		interviewRequestResponseDTOs = append(interviewRequestResponseDTOs, interviewRequestResponseDTO)
	}

	json.NewEncoder(w).Encode(interviewRequestResponseDTOs)
}

func (handler *InterviewRequestHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var interviewRequestResponseDTO model.InterviewRequestResponseDTO
	interviewRequest.FromModel(&interviewRequestResponseDTO)

	json.NewEncoder(w).Encode(interviewRequestResponseDTO)
}

func (handler *InterviewRequestHandler) GetAllBySenderID(w http.ResponseWriter, r *http.Request) {
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

	var interviewRequestResponseDTOs []model.InterviewRequestResponseDTO
	for _, interviewRequest := range *interviewRequests {
		var interviewRequestResponseDTO model.InterviewRequestResponseDTO
		interviewRequest.FromModel(&interviewRequestResponseDTO)
		interviewRequestResponseDTOs = append(interviewRequestResponseDTOs, interviewRequestResponseDTO)
	}

	json.NewEncoder(w).Encode(interviewRequests)
}

func (handler *InterviewRequestHandler) GetAllByReceiverID(w http.ResponseWriter, r *http.Request) {
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

	var interviewRequestResponseDTOs []model.InterviewRequestResponseDTO
	for _, interviewRequest := range *interviewRequests {
		var interviewRequestResponseDTO model.InterviewRequestResponseDTO
		interviewRequest.FromModel(&interviewRequestResponseDTO)
		interviewRequestResponseDTOs = append(interviewRequestResponseDTOs, interviewRequestResponseDTO)
	}

	json.NewEncoder(w).Encode(interviewRequests)
}

func (handler *InterviewRequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var interviewRequestDTO model.InterviewRequestCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&interviewRequestDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	interviewRequest := &model.InterviewRequest{}
	interviewRequest.FromDTO(&interviewRequestDTO)

	err := handler.InterviewRequestService.Create(interviewRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *InterviewRequestHandler) Update(w http.ResponseWriter, r *http.Request) {
	var interviewRequestDTO model.InterviewRequestUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&interviewRequestDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	interviewRequest := &model.InterviewRequest{}
	interviewRequest.FromUpdateDTO(&interviewRequestDTO)

	err := handler.InterviewRequestService.Update(interviewRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
