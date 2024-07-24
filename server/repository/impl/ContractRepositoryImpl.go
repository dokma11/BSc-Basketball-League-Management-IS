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

		if option == "PLAYER_OPTION" {
			contract.OpcUgo = 0
		} else if option == "TEAM_OPTION" {
			contract.OpcUgo = 1
		} else if option == "NO_OPTION" {
			contract.OpcUgo = 2
		}

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

	if option == "PLAYER_OPTION" {
		contract.OpcUgo = 0
	} else if option == "TEAM_OPTION" {
		contract.OpcUgo = 1
	} else if option == "NO_OPTION" {
		contract.OpcUgo = 2
	}

	return &contract, nil
}
