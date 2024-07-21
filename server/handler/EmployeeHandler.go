package handler

import (
	"basketball-league-server/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeHandler(EmployeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: EmployeeService}
}

func (handler *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	employees, err := handler.EmployeeService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employees) // Proveriti samo da li valja
}

func (handler *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) { // Ovde proveriti da li su neophodni parametri
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := handler.EmployeeService.GetByID(id)
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
