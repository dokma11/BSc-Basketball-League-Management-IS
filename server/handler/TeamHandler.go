package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	TeamService *service.TeamService
}

func NewTeamHandler(TeamService *service.TeamService) *TeamHandler {
	return &TeamHandler{TeamService: TeamService}
}

func (handler *TeamHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	teams, err := handler.TeamService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var teamResponseDTOs []model.TeamResponseDTO
	for _, team := range *teams {
		var teamResponseDTO model.TeamResponseDTO
		team.FromModel(&teamResponseDTO)
		teamResponseDTOs = append(teamResponseDTOs, teamResponseDTO)
	}

	json.NewEncoder(w).Encode(teamResponseDTOs)
}

func (handler *TeamHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	team, err := handler.TeamService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.NotFound(w, r)
		return
	}

	var teamResponseDTO model.TeamResponseDTO
	team.FromModel(&teamResponseDTO)
	json.NewEncoder(w).Encode(teamResponseDTO)
}

func (handler *TeamHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	team, err := handler.TeamService.GetByUserID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.NotFound(w, r)
		return
	}

	var teamResponseDTO model.TeamResponseDTO
	team.FromModel(&teamResponseDTO)
	json.NewEncoder(w).Encode(teamResponseDTO)
}

func (handler *TeamHandler) GetPlayerTradeDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeSubjectID, err := strconv.Atoi(vars["tradeSubjectId"])
	if err != nil {
		http.Error(w, "Invalid trade subject ID", http.StatusBadRequest)
		return
	}

	team, err := handler.TeamService.GetPlayerTradeDestination(tradeSubjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.NotFound(w, r)
		return
	}

	var teamResponseDTO model.TeamResponseDTO
	team.FromModel(&teamResponseDTO)
	json.NewEncoder(w).Encode(teamResponseDTO)
}

func (handler *TeamHandler) GetPickTradeDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeSubjectID, err := strconv.Atoi(vars["tradeSubjectId"])
	if err != nil {
		http.Error(w, "Invalid trade subject ID", http.StatusBadRequest)
		return
	}

	team, err := handler.TeamService.GetPickTradeDestination(tradeSubjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.NotFound(w, r)
		return
	}

	var teamResponseDTO model.TeamResponseDTO
	team.FromModel(&teamResponseDTO)
	json.NewEncoder(w).Encode(teamResponseDTO)
}

func (handler *TeamHandler) GetDraftRightsTradeDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeSubjectID, err := strconv.Atoi(vars["tradeSubjectId"])
	if err != nil {
		http.Error(w, "Invalid trade subject ID", http.StatusBadRequest)
		return
	}

	team, err := handler.TeamService.GetDraftRightsTradeDestination(tradeSubjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.NotFound(w, r)
		return
	}

	var teamResponseDTO model.TeamResponseDTO
	team.FromModel(&teamResponseDTO)
	json.NewEncoder(w).Encode(teamResponseDTO)
}

func (handler *TeamHandler) GetWishlistByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	wishlist, err := handler.TeamService.GetWishlistByTeamID(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wishlist) == 0 {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(wishlist)
}
