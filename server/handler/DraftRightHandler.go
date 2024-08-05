package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DraftRightHandler struct {
	DraftRightService *service.DraftRightService
}

func NewDraftRightHandler(DraftRightService *service.DraftRightService) *DraftRightHandler {
	return &DraftRightHandler{DraftRightService: DraftRightService}
}

func (handler *DraftRightHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	draftRights, err := handler.DraftRightService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var draftRightsResponseDTOs []model.DraftRightResponseDTO
	for _, draftRight := range *draftRights {
		var draftRightsResponseDTO model.DraftRightResponseDTO
		draftRight.FromModel(&draftRightsResponseDTO)
		draftRightsResponseDTOs = append(draftRightsResponseDTOs, draftRightsResponseDTO)
	}

	json.NewEncoder(w).Encode(draftRightsResponseDTOs)
}

func (handler *DraftRightHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	draftRight, err := handler.DraftRightService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if draftRight == nil {
		http.NotFound(w, r)
		return
	}

	var draftRightsResponseDTO model.DraftRightResponseDTO
	draftRight.FromModel(&draftRightsResponseDTO)

	json.NewEncoder(w).Encode(draftRightsResponseDTO)
}

func (handler *DraftRightHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	draftRights, err := handler.DraftRightService.GetAllByTeamID(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var draftRightsResponseDTOs []model.DraftRightResponseDTO
	for _, draftRight := range draftRights {
		var draftRightsResponseDTO model.DraftRightResponseDTO
		draftRight.FromModel(&draftRightsResponseDTO)
		draftRightsResponseDTOs = append(draftRightsResponseDTOs, draftRightsResponseDTO)
	}

	json.NewEncoder(w).Encode(draftRightsResponseDTOs)
}

func (handler *DraftRightHandler) GetAllAvailableByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	draftRights, err := handler.DraftRightService.GetAllAvailableByTeamID(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var draftRightsResponseDTOs []model.DraftRightResponseDTO
	for _, draftRight := range draftRights {
		var draftRightsResponseDTO model.DraftRightResponseDTO
		draftRight.FromModel(&draftRightsResponseDTO)
		draftRightsResponseDTOs = append(draftRightsResponseDTOs, draftRightsResponseDTO)
	}

	json.NewEncoder(w).Encode(draftRightsResponseDTOs)
}

func (handler *DraftRightHandler) Update(w http.ResponseWriter, r *http.Request) {
	var draftRightsDTO model.DraftRightUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&draftRightsDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	draftRights := &model.DraftRight{}
	draftRights.FromUpdateDTO(&draftRightsDTO)

	err := handler.DraftRightService.Update(draftRights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *DraftRightHandler) AddToWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var draftRightsDTO model.DraftRightCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&draftRightsDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	draftRights := &model.DraftRight{}
	draftRights.FromCreateDTO(&draftRightsDTO)

	err = handler.DraftRightService.AddToWishlist(draftRights, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *DraftRightHandler) RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var draftRightsDTO model.DraftRightCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&draftRightsDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	draftRights := &model.DraftRight{}
	draftRights.FromCreateDTO(&draftRightsDTO)

	err = handler.DraftRightService.RemoveFromWishlist(draftRights, teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
