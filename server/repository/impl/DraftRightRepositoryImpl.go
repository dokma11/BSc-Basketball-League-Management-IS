package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type draftRightRepository struct {
	db *sql.DB
}

func NewDraftRightRepository(db *sql.DB) repository.DraftRightRepository {
	return &draftRightRepository{db}
}

func (repo *draftRightRepository) GetAll() ([]model.DraftRight, error) {
	rows, err := repo.db.Query("SELECT * FROM PravaNaIgraca") // Proveriti samo naziv
	if err != nil {
		return nil, fmt.Errorf("failed to query all draft rights: %v", err)
	}
	defer rows.Close()

	var draftRights []model.DraftRight
	for rows.Next() {
		var draftRight model.DraftRight
		var position string
		if err := rows.Scan(&draftRight.IdPrava, &draftRight.ImeIgrPrava, &draftRight.PrezimeIgrPrava, &position,
			&draftRight.IdTim, &draftRight.IdRegrut, &draftRight.IdPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if position == "PG" {
			draftRight.PozicijaIgrPrava = 1
		} else if position == "SG" {
			draftRight.PozicijaIgrPrava = 2
		} else if position == "SF" {
			draftRight.PozicijaIgrPrava = 3
		} else if position == "PF" {
			draftRight.PozicijaIgrPrava = 4
		} else if position == "C" {
			draftRight.PozicijaIgrPrava = 5
		}

		draftRights = append(draftRights, draftRight)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return draftRights, nil
}

func (repo *draftRightRepository) GetByID(id int) (*model.DraftRight, error) {
	var draftRight model.DraftRight
	var position string
	row := repo.db.QueryRow("SELECT * FROM PravaNaIgraca WHERE IDPRAVA = :1", id) // Proveriti samo naziv
	if err := row.Scan(&draftRight.IdPrava, &draftRight.ImeIgrPrava, &draftRight.PrezimeIgrPrava, &position,
		&draftRight.IdTim, &draftRight.IdRegrut, &draftRight.IdPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	if position == "PG" {
		draftRight.PozicijaIgrPrava = 1
	} else if position == "SG" {
		draftRight.PozicijaIgrPrava = 2
	} else if position == "SF" {
		draftRight.PozicijaIgrPrava = 3
	} else if position == "PF" {
		draftRight.PozicijaIgrPrava = 4
	} else if position == "C" {
		draftRight.PozicijaIgrPrava = 5
	}

	return &draftRight, nil
}
