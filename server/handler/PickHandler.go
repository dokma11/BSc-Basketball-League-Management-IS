package handler

import (
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

func (handler *PickHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	picks, err := handler.PickService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}

func (handler *PickHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
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

	json.NewEncoder(w).Encode(pick)
}

func (handler *PickHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
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

	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}

func (handler *PickHandler) GetAllByYear(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	year, _ := vars["year"]
	picks, _ := handler.PickService.GetAllByYear(year)
	json.NewEncoder(w).Encode(picks) // Proveriti samo da li valja
}
