package handler

import (
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

	json.NewEncoder(w).Encode(teams)
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

	json.NewEncoder(w).Encode(team)
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

	json.NewEncoder(w).Encode(team)
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

	json.NewEncoder(w).Encode(team)
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

	json.NewEncoder(w).Encode(team)
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

	json.NewEncoder(w).Encode(team)
}
