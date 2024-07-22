package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewEmployeeService(EmployeeRepository repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{EmployeeRepository: EmployeeRepository}
}

func (service *EmployeeService) GetAll() (*[]model.Employee, error) {
	employees, err := service.EmployeeRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no employees were found"))
	}

	return &employees, nil
}

func (service *EmployeeService) GetByID(id int) (*model.Employee, error) {
	employee, err := service.EmployeeRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no employees with that id were found"))
	}

	return employee, nil
}
