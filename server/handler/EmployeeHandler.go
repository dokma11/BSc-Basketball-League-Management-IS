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

type EmployeeHandler struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeHandler(EmployeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: EmployeeService}
}

func (handler *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := handler.EmployeeService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var employeeResponseDTOs []model.EmployeeResponseDTO
	for _, employee := range *employees {
		var employeeResponseDTO model.EmployeeResponseDTO
		employee.FromModel(&employeeResponseDTO)
		employeeResponseDTOs = append(employeeResponseDTOs, employeeResponseDTO)
	}

	json.NewEncoder(w).Encode(employeeResponseDTOs)
}

func (handler *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	var employeeResponseDTO model.EmployeeResponseDTO
	employee.FromModel(&employeeResponseDTO)

	json.NewEncoder(w).Encode(employeeResponseDTO)
}

func (handler *EmployeeHandler) GetByTeamID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	employee, err := handler.EmployeeService.GetByTeamID(teamID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if employee == nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	var employeeResponseDTO model.EmployeeResponseDTO
	employee.FromModel(&employeeResponseDTO)

	json.NewEncoder(w).Encode(employeeResponseDTO)
}
