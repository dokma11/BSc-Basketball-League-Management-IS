package handler

import (
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

func (handler *DraftHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	drafts, err := handler.DraftService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(drafts) // Proveriti samo da li valja
}

func (handler *DraftHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
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

	json.NewEncoder(w).Encode(draft)
}
