package repository

import "basketball-league-server/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}
