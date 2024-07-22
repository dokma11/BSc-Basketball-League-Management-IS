package repository

import "basketball-league-server/model"

type InterviewRepository interface {
	GetAll() ([]model.Interview, error)
	GetByID(id int) (*model.Interview, error)
	GetAllByUserID(userID int) ([]model.Interview, error)
	Create(*model.Interview) error
}
