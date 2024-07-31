package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DraftHandler struct {
	DraftService *service.DraftService
}

func NewDraftHandler(DraftService *service.DraftService) *DraftHandler {
	return &DraftHandler{DraftService: DraftService}
}

func (handler *DraftHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	drafts, err := handler.DraftService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var draftResponseDTOs []model.DraftResponseDTO
	for _, draft := range *drafts {
		var draftResponseDTO model.DraftResponseDTO
		draft.FromModel(&draftResponseDTO)
		draftResponseDTOs = append(draftResponseDTOs, draftResponseDTO)
	}

	json.NewEncoder(w).Encode(draftResponseDTOs)
}

func (handler *DraftHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	draft, err := handler.DraftService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if draft == nil {
		http.NotFound(w, r)
		return
	}

	var draftResponseDTO model.DraftResponseDTO
	draft.FromModel(&draftResponseDTO)

	json.NewEncoder(w).Encode(draftResponseDTO)
}
