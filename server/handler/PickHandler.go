package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PickHandler struct {
	PickService *service.PickService
}

func NewPickHandler(PickService *service.PickService) *PickHandler {
	return &PickHandler{PickService: PickService}
}

func (handler *PickHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	picks, err := handler.PickService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pickResponseDTOs []model.PickResponseDTO
	for _, pick := range *picks {
		var pickResponseDTO model.PickResponseDTO
		pick.FromModel(&pickResponseDTO)
		pickResponseDTOs = append(pickResponseDTOs, pickResponseDTO)
	}

	json.NewEncoder(w).Encode(pickResponseDTOs)
}

func (handler *PickHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pick, err := handler.PickService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pick == nil {
		http.NotFound(w, r)
		return
	}

	var pickResponseDTO model.PickResponseDTO
	pick.FromModel(&pickResponseDTO)

	json.NewEncoder(w).Encode(pickResponseDTO)
}

func (handler *PickHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	picks, err := handler.PickService.GetAllByTeamID(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pickResponseDTOs []model.PickResponseDTO
	for _, pick := range *picks {
		var pickResponseDTO model.PickResponseDTO
		pick.FromModel(&pickResponseDTO)
		pickResponseDTOs = append(pickResponseDTOs, pickResponseDTO)
	}

	json.NewEncoder(w).Encode(pickResponseDTOs)
}

func (handler *PickHandler) GetAllAvailableByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	picks, err := handler.PickService.GetAllAvailableByTeamID(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pickResponseDTOs []model.PickResponseDTO
	for _, pick := range *picks {
		var pickResponseDTO model.PickResponseDTO
		pick.FromModel(&pickResponseDTO)
		pickResponseDTOs = append(pickResponseDTOs, pickResponseDTO)
	}

	json.NewEncoder(w).Encode(pickResponseDTOs)
}

func (handler *PickHandler) GetAllByYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, _ := vars["year"]
	picks, _ := handler.PickService.GetAllByYear(year)

	var pickResponseDTOs []model.PickResponseDTO
	for _, pick := range *picks {
		var pickResponseDTO model.PickResponseDTO
		pick.FromModel(&pickResponseDTO)
		pickResponseDTOs = append(pickResponseDTOs, pickResponseDTO)
	}

	json.NewEncoder(w).Encode(pickResponseDTOs)
}

func (handler *PickHandler) Update(w http.ResponseWriter, r *http.Request) {
	var pickDTO model.PickUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&pickDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pick := &model.Pick{}
	pick.FromUpdateDTO(&pickDTO)

	err := handler.PickService.Update(pick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *PickHandler) AddToWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var pickDTO model.PickCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&pickDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pick := &model.Pick{}
	pick.FromCreateDTO(&pickDTO)

	err = handler.PickService.AddToWishlist(pick, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *PickHandler) RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var pickDTO model.PickCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&pickDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pick := &model.Pick{}
	pick.FromCreateDTO(&pickDTO)

	err = handler.PickService.RemoveFromWishlist(pick, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
