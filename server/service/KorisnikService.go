package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type KorisnikService struct {
	KorisnikRepository repository.KorisnikRepository
}

func NewKorisnikService(korisnikRepository repository.KorisnikRepository) *KorisnikService {
	return &KorisnikService{KorisnikRepository: korisnikRepository}
}

func (service *KorisnikService) GetAll() (*[]model.Korisnik, error) {
	users, err := service.KorisnikRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users were found"))
	}

	return &users, nil
}

func (service *KorisnikService) GetByID(id int) (*model.Korisnik, error) {
	user, err := service.KorisnikRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with that id were found"))
	}

	return user, nil
}
