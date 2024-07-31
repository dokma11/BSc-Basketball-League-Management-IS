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
