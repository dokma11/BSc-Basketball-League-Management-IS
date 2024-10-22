package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
	"log"
)

type ContractService struct {
	ContractRepository repository.ContractRepository
}

func NewContractService(ContractRepository repository.ContractRepository) *ContractService {
	return &ContractService{ContractRepository: ContractRepository}
}

func (service *ContractService) GetAll() (*[]model.Contract, error) {
	contracts, err := service.ContractRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no contracts were found"))
	}
	return &contracts, nil
}

func (service *ContractService) GetByID(id int) (*model.Contract, error) {
	contract, err := service.ContractRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no contracts with that id were found"))
	}
	return contract, nil
}

func (service *ContractService) Update(contract *model.Contract) error {
	err := service.ContractRepository.Update(contract)
	if err != nil {
		log.Println("Error updating contract")
		return err
	}
	return nil
}
