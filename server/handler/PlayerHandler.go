package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PlayerHandler struct {
	PlayerService *service.PlayerService
}

func NewPlayerHandler(PlayerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{PlayerService: PlayerService}
}

func (handler *PlayerHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	players, err := handler.PlayerService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players) // Proveriti samo da li valja
}

func (handler *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	player, err := handler.PlayerService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if player == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(player)
}

func (handler *PlayerHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	players, err := handler.PlayerService.GetAllByTeamID(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players) // Proveriti samo da li valja
}
