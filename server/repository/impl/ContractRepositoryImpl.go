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
		if err := rows.Scan(&contract.IdUgo, &contract.DatPotUgo, &contract.DatVazUgo,
			&contract.VredUgo, &contract.OpcUgo); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
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
	row := repo.db.QueryRow("SELECT * FROM UGOVOR WHERE IDUGO = :1", id)
	if err := row.Scan(&contract.IdUgo, &contract.DatPotUgo, &contract.DatVazUgo,
		&contract.VredUgo, &contract.OpcUgo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &contract, nil
}
