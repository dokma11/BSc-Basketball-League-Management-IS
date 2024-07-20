package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PikHandler struct {
	PikService *service.PikService
}

func NewPikHandler(pikService *service.PikService) *PikHandler {
	return &PikHandler{PikService: pikService}
}

func (handler *PikHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	picks, err := handler.PikService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}

func (handler *PikHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pick, err := handler.PikService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pick == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(pick)
}

func (handler *PikHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	picks, err := handler.PikService.GetAllByTeamID(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}

func (handler *PikHandler) GetAllByYear(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	year, _ := vars["year"]
	picks, _ := handler.PikService.GetAllByYear(year)
	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}
