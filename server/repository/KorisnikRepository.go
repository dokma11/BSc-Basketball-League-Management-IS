package repository

import "basketball-league-server/model"

type KorisnikRepository interface {
	GetAll() ([]model.Korisnik, error)
	GetByID(id int) (*model.Korisnik, error)
	GetByEmail(email string) (*model.Korisnik, error)
}
