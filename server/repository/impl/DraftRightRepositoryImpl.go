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
	rows, err := repo.db.Query("SELECT * FROM PravaNaIgraca")
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

		mapDraftRightsEnum(position, &draftRight)

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
	row := repo.db.QueryRow("SELECT * FROM PravaNaIgraca WHERE IDPRAVA = :1", id)
	if err := row.Scan(&draftRight.IdPrava, &draftRight.ImeIgrPrava, &draftRight.PrezimeIgrPrava, &position,
		&draftRight.IdTim, &draftRight.IdRegrut, &draftRight.IdPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	mapDraftRightsEnum(position, &draftRight)

	return &draftRight, nil
}

func (repo *draftRightRepository) GetAllByTeamID(teamID int) ([]model.DraftRight, error) {
	rows, err := repo.db.Query("SELECT * FROM PravaNaIgraca WHERE IDTIM = :1", teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
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

		mapDraftRightsEnum(position, &draftRight)

		draftRights = append(draftRights, draftRight)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return draftRights, nil
}

func (repo *draftRightRepository) Update(draftRights *model.DraftRight) error {
	_, err := repo.db.Exec("UPDATE PravaNaIgraca SET IDTIM = :1 WHERE IDPRAVA = :2", draftRights.IdTim, draftRights.IdPrava)
	if err != nil {
		return fmt.Errorf("failed to update draft rights: %v", err)
	}
	return nil
}

func mapDraftRightsEnum(position string, draftRights *model.DraftRight) {
	switch position {
	case "PG":
		draftRights.PozicijaIgrPrava = 0
	case "SG":
		draftRights.PozicijaIgrPrava = 1
	case "SF":
		draftRights.PozicijaIgrPrava = 2
	case "PF":
		draftRights.PozicijaIgrPrava = 3
	default:
		draftRights.PozicijaIgrPrava = 4
	}
}
