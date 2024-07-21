package handler

import (
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

func (handler *DraftRightHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	draftRights, err := handler.DraftRightService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(draftRights) // Proveriti samo da li valja
}

func (handler *DraftRightHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
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

	json.NewEncoder(w).Encode(draftRight)
}
