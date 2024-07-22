package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TradeHandler struct {
	TradeService *service.TradeService
}

func NewTradeHandler(TradeService *service.TradeService) *TradeHandler {
	return &TradeHandler{TradeService: TradeService}
}

func (handler *TradeHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	trades, err := handler.TradeService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trades) // Proveriti samo da li valja
}

func (handler *TradeHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trade, err := handler.TradeService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if trade == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(trade)
}

func (handler *TradeHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	trades, err := handler.TradeService.GetAllByTeamID(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trades) // Proveriti samo da li valja
}

func (handler *TradeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var trade model.Trade
	if err := json.NewDecoder(r.Body).Decode(&trade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TradeService.Create(&trade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
