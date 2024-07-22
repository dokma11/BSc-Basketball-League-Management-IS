package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TradeProposalHandler struct {
	TradeProposalService *service.TradeProposalService
}

func NewTradeProposalHandler(TradeProposalService *service.TradeProposalService) *TradeProposalHandler {
	return &TradeProposalHandler{TradeProposalService: TradeProposalService}
}

func (handler *TradeProposalHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	tradeProposals, err := handler.TradeProposalService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tradeProposals) // Proveriti samo da li valja
}

func (handler *TradeProposalHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tradeProposal, err := handler.TradeProposalService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tradeProposal == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(tradeProposal)
}

func (handler *TradeProposalHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tradeProposals, err := handler.TradeProposalService.GetAllByTeamID(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tradeProposals) // Proveriti samo da li valja
}

func (handler *TradeProposalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tradeProposal model.TradeProposal
	if err := json.NewDecoder(r.Body).Decode(&tradeProposal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TradeProposalService.Create(&tradeProposal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TradeProposalHandler) Update(w http.ResponseWriter, r *http.Request) {
	var tradeProposal model.TradeProposal
	if err := json.NewDecoder(r.Body).Decode(&tradeProposal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.TradeProposalService.Update(&tradeProposal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
