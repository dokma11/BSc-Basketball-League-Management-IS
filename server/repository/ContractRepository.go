package repository

import "basketball-league-server/model"

type ContractRepository interface {
	GetAll() ([]model.Contract, error)
	GetByID(id int) (*model.Contract, error)
	Update(contract *model.Contract) error
}
