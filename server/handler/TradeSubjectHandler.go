package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TradeSubjectHandler struct {
	TradeSubjectService  *service.TradeSubjectService
	TradeProposalService *service.TradeProposalService
}

func NewTradeSubjectHandler(TradeSubjectService *service.TradeSubjectService, TradeProposalService *service.TradeProposalService) *TradeSubjectHandler {
	return &TradeSubjectHandler{TradeSubjectService: TradeSubjectService, TradeProposalService: TradeProposalService}
}

func (handler *TradeSubjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tradeSubjects, err := handler.TradeSubjectService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tradeSubjects)
}

func (handler *TradeSubjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tradeSubject, err := handler.TradeSubjectService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tradeSubject == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(tradeSubject)
}

func (handler *TradeSubjectHandler) GetAllByTradeID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeID, err := strconv.Atoi(vars["tradeId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tradeSubjects, err := handler.TradeSubjectService.GetAllByTradeID(tradeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tradeSubjects)
}

func (handler *TradeSubjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tradeSubjectDTO model.TradeSubjectCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&tradeSubjectDTO); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tradeSubject, tradeSubjectError := handler.mapSubjectFromDTO(&tradeSubjectDTO)
	if tradeSubjectError != nil {
		log.Println(tradeSubjectError)
		http.Error(w, tradeSubjectError.Error(), http.StatusInternalServerError)
		return
	}

	err := handler.TradeSubjectService.Create(tradeSubject)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TradeSubjectHandler) mapSubjectFromDTO(tradeSubjectDTO *model.TradeSubjectCreateDTO) (*model.TradeSubject, error) {
	var tradeSubject model.TradeSubject
	tradeSubject.IdPik = tradeSubjectDTO.IdPik
	tradeSubject.IdPrava = tradeSubjectDTO.IdPrava
	tradeSubject.IdIgrac = tradeSubjectDTO.IdIgrac

	if tradeSubjectDTO.TipPredTrg == 0 {
		tradeSubject.TipPredTrg = 0
	} else if tradeSubjectDTO.TipPredTrg == 1 {
		tradeSubject.TipPredTrg = 1
	} else if tradeSubjectDTO.TipPredTrg == 2 {
		tradeSubject.TipPredTrg = 2
	}

	latestTradeProposal, err := handler.TradeProposalService.GetLatest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tradeSubject.IdZahTrg = latestTradeProposal.IdZahTrg

	return &tradeSubject, nil
}
