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

func (handler *TeamHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	teams, err := handler.TeamService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams) // Proveriti samo da li valja
}

func (handler *TeamHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
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
