package repository

import "basketball-league-server/model"

type RecruitRepository interface {
	GetAll() ([]model.Recruit, error)
	GetByID(id int) (*model.Recruit, error)
	Create(*model.Recruit) error
	Update(*model.Recruit) error
	AddToWishlist(recruit *model.Recruit, teamId int) error
	RemoveFromWishlist(recruit *model.Recruit, teamId int) error
	GetAllByName(name string) ([]model.Recruit, error)
}
