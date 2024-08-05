package handler

import (
	"basketball-league-server/model"
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

func (handler *PlayerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	players, err := handler.PlayerService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var playerResponseDTOs []model.PlayerResponseDTO
	for _, player := range *players {
		var playerResponseDTO model.PlayerResponseDTO
		player.FromModel(&playerResponseDTO)
		playerResponseDTOs = append(playerResponseDTOs, playerResponseDTO)
	}

	json.NewEncoder(w).Encode(playerResponseDTOs)
}

func (handler *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var playerResponseDTO model.PlayerResponseDTO
	player.FromModel(&playerResponseDTO)

	json.NewEncoder(w).Encode(playerResponseDTO)
}

func (handler *PlayerHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) {
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

	var playerResponseDTOs []model.PlayerResponseDTO
	for _, player := range *players {
		var playerResponseDTO model.PlayerResponseDTO
		player.FromModel(&playerResponseDTO)
		playerResponseDTOs = append(playerResponseDTOs, playerResponseDTO)
	}

	json.NewEncoder(w).Encode(playerResponseDTOs)
}

func (handler *PlayerHandler) GetAllAvailableByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	players, err := handler.PlayerService.GetAllAvailableByTeamID(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var playerResponseDTOs []model.PlayerResponseDTO
	for _, player := range *players {
		var playerResponseDTO model.PlayerResponseDTO
		player.FromModel(&playerResponseDTO)
		playerResponseDTOs = append(playerResponseDTOs, playerResponseDTO)
	}

	json.NewEncoder(w).Encode(playerResponseDTOs)
}

func (handler *PlayerHandler) Update(w http.ResponseWriter, r *http.Request) {
	var playerDTO model.PlayerUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&playerDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player := &model.Player{}
	player.FromUpdateDTO(&playerDTO)

	err := handler.PlayerService.Update(player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
