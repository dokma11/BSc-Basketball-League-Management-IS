package handler

import (
	"basketball-league-server/model"
	"basketball-league-server/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TradeSubjectHandler struct {
	TradeSubjectService  *service.TradeSubjectService
	TradeProposalService *service.TradeProposalService
	TeamService          *service.TeamService
	PickService          *service.PickService
	DraftRightService    *service.DraftRightService
	EmployeeService      *service.EmployeeService
	ContractService      *service.ContractService
	TradeService         *service.TradeService
}

func NewTradeSubjectHandler(TradeSubjectService *service.TradeSubjectService, TradeProposalService *service.TradeProposalService,
	TeamService *service.TeamService, PickService *service.PickService, DraftRightService *service.DraftRightService,
	EmployeeService *service.EmployeeService, ContractService *service.ContractService, TradeService *service.TradeService) *TradeSubjectHandler {
	return &TradeSubjectHandler{TradeSubjectService: TradeSubjectService, TradeProposalService: TradeProposalService,
		TeamService: TeamService, PickService: PickService, DraftRightService: DraftRightService, EmployeeService: EmployeeService,
		ContractService: ContractService, TradeService: TradeService}
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

	tradeSubjects, err := handler.TradeSubjectService.GetAllByTradeProposalID(tradeID)
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

func (handler *TradeSubjectHandler) CommitTrade(w http.ResponseWriter, r *http.Request) {
	var tradeProposalDTO model.TradeProposalUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&tradeProposalDTO); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tradeSubjects, err := handler.TradeSubjectService.GetAllByTradeProposalID(int(tradeProposalDTO.IdZahTrg))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, tradeSubject := range *tradeSubjects {
		if tradeSubject.Type == 0 { // Player
			team, err := handler.TeamService.GetPlayerTradeDestination(int(tradeSubject.ID))
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			player, err := handler.EmployeeService.GetByID(int(tradeSubject.PlayerId)) // Player is considered as an employee
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			contract, err := handler.ContractService.GetByID(int(player.ContractId))
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			contract.TeamId = team.ID // Switch player's team
			contractError := handler.ContractService.Update(contract)
			if contractError != nil {
				log.Println(contractError)
				return
			}
		} else if tradeSubject.Type == 1 { // Pick
			team, err := handler.TeamService.GetPickTradeDestination(int(tradeSubject.ID))
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			pick, err := handler.PickService.GetByID(int(tradeSubject.PickId))
			if err != nil {
				log.Println(err)
			}

			pick.TeamId = team.ID
			pickError := handler.PickService.Update(pick)
			if pickError != nil {
				log.Println(pickError)
			}
		} else if tradeSubject.Type == 2 { // Draft Right
			team, err := handler.TeamService.GetDraftRightsTradeDestination(int(tradeSubject.ID))
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			draftRights, err := handler.DraftRightService.GetByID(int(tradeSubject.DraftRightsId))
			if err != nil {
				log.Println(err)
			}

			draftRights.TeamId = team.ID
			draftRightsError := handler.DraftRightService.Update(draftRights)
			if draftRightsError != nil {
				log.Println(draftRightsError)
			}
		}
	}

	var trade, tradeErr = model.NewTrade(time.Now(), tradeProposalDTO.TipZahTrg, tradeProposalDTO.IdZahTrg)
	if tradeErr != nil {
		log.Println(tradeErr)
	}

	tradeError := handler.TradeService.Create(trade)
	if tradeError != nil {
		fmt.Println(tradeError)
	}
}

func (handler *TradeSubjectHandler) GetDetailsForTradeProposal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeProposalID, err := strconv.Atoi(vars["tradeId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	playerTypeTradeSubjects, err := handler.TradeSubjectService.GetPlayerTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pickTypeTradeSubjects, err := handler.TradeSubjectService.GetPickTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	draftRightsTypeTradeSubjects, err := handler.TradeSubjectService.GetDraftRightsTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var tradeSubjectsDTO []model.TradeSubjectDetailsResponseDTO
	if len(*playerTypeTradeSubjects) > 0 {
		for _, tradeSubject := range *playerTypeTradeSubjects {
			tradeSubjectsDTO = append(tradeSubjectsDTO, tradeSubject)
		}
	}
	if len(*pickTypeTradeSubjects) > 0 {
		for _, tradeSubject := range *pickTypeTradeSubjects {
			tradeSubjectsDTO = append(tradeSubjectsDTO, tradeSubject)
		}
	}
	if len(*draftRightsTypeTradeSubjects) > 0 {
		for _, tradeSubject := range *draftRightsTypeTradeSubjects {
			tradeSubjectsDTO = append(tradeSubjectsDTO, tradeSubject)
		}
	}

	json.NewEncoder(w).Encode(tradeSubjectsDTO)
}

func (handler *TradeSubjectHandler) mapSubjectFromDTO(tradeSubjectDTO *model.TradeSubjectCreateDTO) (*model.TradeSubject, error) {
	var tradeSubject model.TradeSubject
	tradeSubject.PickId = tradeSubjectDTO.IdPik
	tradeSubject.DraftRightsId = tradeSubjectDTO.IdPrava
	tradeSubject.PlayerId = tradeSubjectDTO.IdIgrac

	if tradeSubjectDTO.TipPredTrg == 0 {
		tradeSubject.Type = 0
	} else if tradeSubjectDTO.TipPredTrg == 1 {
		tradeSubject.Type = 1
	} else if tradeSubjectDTO.TipPredTrg == 2 {
		tradeSubject.Type = 2
	}

	latestTradeProposal, err := handler.TradeProposalService.GetLatest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tradeSubject.TradeProposalId = latestTradeProposal.ID

	return &tradeSubject, nil
}
