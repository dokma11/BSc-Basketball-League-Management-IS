package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type contractRepository struct {
	db *sql.DB
}

func NewContractRepository(db *sql.DB) repository.ContractRepository {
	return &contractRepository{db}
}

func (repo *contractRepository) GetAll() ([]model.Contract, error) {
	rows, err := repo.db.Query("SELECT * FROM UGOVOR")
	if err != nil {
		return nil, fmt.Errorf("failed to query all Contracts: %v", err)
	}
	defer rows.Close()

	var contracts []model.Contract
	for rows.Next() {
		var contract model.Contract
		var option string
		if err := rows.Scan(&contract.IdUgo, &contract.DatPotUgo, &contract.DatVazUgo, &contract.VredUgo,
			&option, &contract.IdTim, &contract.IdTipUgo); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		mapContractEnum(option, &contract)

		contracts = append(contracts, contract)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return contracts, nil
}

func (repo *contractRepository) GetByID(id int) (*model.Contract, error) {
	var contract model.Contract
	var option string
	row := repo.db.QueryRow("SELECT * FROM UGOVOR WHERE IDUGO = :1", id)
	if err := row.Scan(&contract.IdUgo, &contract.DatPotUgo, &contract.DatVazUgo, &contract.VredUgo,
		&option, &contract.IdTim, &contract.IdTipUgo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	mapContractEnum(option, &contract)

	return &contract, nil
}

func (repo *contractRepository) Update(contract *model.Contract) error {
	_, err := repo.db.Exec("UPDATE UGOVOR SET IDTIM = :1 WHERE IDUGO = :2", contract.IdTim, contract.IdUgo)
	if err != nil {
		return fmt.Errorf("failed to update contract: %v", err)
	}
	return nil
}

func mapContractEnum(option string, contract *model.Contract) {
	switch option {
	case "PLAYER_OPTION":
		contract.OpcUgo = 0
	case "TEAM_OPTION":
		contract.OpcUgo = 1
	default:
		contract.OpcUgo = 2
	}
}
