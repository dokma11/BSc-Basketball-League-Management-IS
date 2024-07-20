package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ZaposleniHandler struct {
	ZaposleniService *service.ZaposleniService
}

func NewZaposleniHandler(zaposleniService *service.ZaposleniService) *ZaposleniHandler {
	return &ZaposleniHandler{ZaposleniService: zaposleniService}
}

func (handler *ZaposleniHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	employees, err := handler.ZaposleniService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employees) // Proveriti samo da li valja
}

func (handler *ZaposleniHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := handler.ZaposleniService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if employee == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(employee)
}
