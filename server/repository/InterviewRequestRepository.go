package repository

import "basketball-league-server/model"

type InterviewRequestRepository interface {
	GetAll() ([]model.InterviewRequest, error)
	GetByID(id int) (*model.InterviewRequest, error)
	GetAllBySenderID(userID int) ([]model.InterviewRequest, error)
	GetAllByReceiverID(userID int) ([]model.InterviewRequest, error)
	Create(*model.InterviewRequest) error
	Update(*model.InterviewRequest) error
}
