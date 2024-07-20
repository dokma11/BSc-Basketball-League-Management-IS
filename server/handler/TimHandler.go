package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TimHandler struct {
	TimService *service.TimService
}

func NewTimHandler(timService *service.TimService) *TimHandler {
	return &TimHandler{TimService: timService}
}

func (handler *TimHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	teams, err := handler.TimService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams) // Proveriti samo da li valja
}

func (handler *TimHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tim, err := handler.TimService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tim == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(tim)
}
