package repository

import "basketball-league-server/model"

type EmployeeRepository interface {
	GetAll() ([]model.Employee, error)
	GetByID(id int) (*model.Employee, error)
}
