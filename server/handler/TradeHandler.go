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

func (handler *TradeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	trades, err := handler.TradeService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tradeResponseDTOs []model.TradeResponseDTO
	for _, trade := range *trades {
		var tradeResponseDTO model.TradeResponseDTO
		trade.FromModel(&tradeResponseDTO)
		tradeResponseDTOs = append(tradeResponseDTOs, tradeResponseDTO)
	}

	json.NewEncoder(w).Encode(tradeResponseDTOs)
}

func (handler *TradeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var tradeResponseDTO model.TradeResponseDTO
	trade.FromModel(&tradeResponseDTO)
	json.NewEncoder(w).Encode(tradeResponseDTO)
}

func (handler *TradeHandler) GetAllByTeamID(w http.ResponseWriter, r *http.Request) {
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

	var tradeResponseDTOs []model.TradeResponseDTO
	for _, trade := range *trades {
		var tradeResponseDTO model.TradeResponseDTO
		trade.FromModel(&tradeResponseDTO)
		tradeResponseDTOs = append(tradeResponseDTOs, tradeResponseDTO)
	}

	json.NewEncoder(w).Encode(tradeResponseDTOs)
}

func (handler *TradeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tradeDTO model.TradeCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&tradeDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	trade := &model.Trade{}
	trade.FromDTO(&tradeDTO)

	err := handler.TradeService.Create(trade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
