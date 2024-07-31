package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ContractHandler struct {
	ContractService *service.ContractService
}

func NewContractHandler(ContractService *service.ContractService) *ContractHandler {
	return &ContractHandler{ContractService: ContractService}
}

func (handler *ContractHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	contracts, err := handler.ContractService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contracts)
}

func (handler *ContractHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	contract, err := handler.ContractService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if contract == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(contract)
}
