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

type TradeProposalHandler struct {
	TradeProposalService *service.TradeProposalService
	EmployeeService      *service.EmployeeService
}

func NewTradeProposalHandler(TradeProposalService *service.TradeProposalService, EmployeeService *service.EmployeeService) *TradeProposalHandler {
	return &TradeProposalHandler{TradeProposalService: TradeProposalService, EmployeeService: EmployeeService}
}

func (handler *TradeProposalHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tradeProposals, err := handler.TradeProposalService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tradeProposalResponseDTOs []model.TradeProposalResponseDTO
	for _, tradeProposal := range *tradeProposals {
		var tradeProposalResponseDTO model.TradeProposalResponseDTO
		tradeProposal.FromModel(&tradeProposalResponseDTO)
		tradeProposalResponseDTOs = append(tradeProposalResponseDTOs, tradeProposalResponseDTO)
	}

	json.NewEncoder(w).Encode(tradeProposalResponseDTOs)
}

func (handler *TradeProposalHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var tradeProposalResponseDTO model.TradeProposalResponseDTO
	tradeProposal.FromModel(&tradeProposalResponseDTO)
	json.NewEncoder(w).Encode(tradeProposalResponseDTO)
}

func (handler *TradeProposalHandler) GetAllReceivedByManagerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	managerID, err := strconv.Atoi(vars["managerId"])
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	tradeProposals, err := handler.TradeProposalService.GetAllReceivedByManagerID(managerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tradeProposalResponseDTOs []model.TradeProposalResponseDTO
	for _, tradeProposal := range *tradeProposals {
		var tradeProposalResponseDTO model.TradeProposalResponseDTO
		tradeProposal.FromModel(&tradeProposalResponseDTO)
		tradeProposalResponseDTOs = append(tradeProposalResponseDTOs, tradeProposalResponseDTO)
	}

	json.NewEncoder(w).Encode(tradeProposalResponseDTOs)
}

func (handler *TradeProposalHandler) GetAllSentByManagerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	managerID, err := strconv.Atoi(vars["managerId"])
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	tradeProposals, err := handler.TradeProposalService.GetAllSentByManagerID(managerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tradeProposalResponseDTOs []model.TradeProposalResponseDTO
	for _, tradeProposal := range *tradeProposals {
		var tradeProposalResponseDTO model.TradeProposalResponseDTO
		tradeProposal.FromModel(&tradeProposalResponseDTO)
		tradeProposalResponseDTOs = append(tradeProposalResponseDTOs, tradeProposalResponseDTO)
	}

	json.NewEncoder(w).Encode(tradeProposalResponseDTOs)
}

func (handler *TradeProposalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tradeProposalDTO model.TradeProposalCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&tradeProposalDTO); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tradeProposal, tradeProposalError := handler.mapFromDTO(&tradeProposalDTO)
	if tradeProposalError != nil {
		log.Println(tradeProposalError)
		http.Error(w, tradeProposalError.Error(), http.StatusInternalServerError)
		return
	}

	err := handler.TradeProposalService.Create(tradeProposal)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TradeProposalHandler) Update(w http.ResponseWriter, r *http.Request) {
	var tradeProposalDTO model.TradeProposalUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&tradeProposalDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tradeProposal := &model.TradeProposal{}
	tradeProposal.FromUpdateDTO(&tradeProposalDTO)

	err := handler.TradeProposalService.Update(tradeProposal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TradeProposalHandler) GetLatest(w http.ResponseWriter, r *http.Request) {
	tradeProposal, err := handler.TradeProposalService.GetLatest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tradeProposal == nil {
		http.NotFound(w, r)
		return
	}

	var tradeProposalResponseDTO model.TradeProposalResponseDTO
	tradeProposal.FromModel(&tradeProposalResponseDTO)
	json.NewEncoder(w).Encode(tradeProposalResponseDTO)
}

func (handler *TradeProposalHandler) mapFromDTO(tradeProposalDTO *model.TradeProposalCreateDTO) (*model.TradeProposal, error) {
	var tradeProposal model.TradeProposal
	tradeProposal.Date = tradeProposalDTO.DatZahTrg
	tradeProposal.Type = tradeProposalDTO.TipZahTrg
	if tradeProposalDTO.TipZahTrg == 0 {
		tradeProposal.Type = 0
	} else if tradeProposalDTO.TipZahTrg == 1 {
		tradeProposal.Type = 1
	} else if tradeProposalDTO.TipZahTrg == 2 {
		tradeProposal.Type = 2
	}

	tradeProposal.Status = 0 // 0 = IN_PROGRESS
	tradeProposal.SenderId = tradeProposalDTO.IdMenadzerPos

	manager, err := handler.EmployeeService.GetByTeamID(int(tradeProposalDTO.IdMenadzerPrimTim))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tradeProposal.ReceiverId = manager.ID
	tradeProposal.DenialReason = ""
	return &tradeProposal, nil
}
