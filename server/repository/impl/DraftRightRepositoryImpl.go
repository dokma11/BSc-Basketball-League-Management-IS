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
	rows, err := repo.db.Query("SELECT * FROM PRAVA_NA_IGRACA") // Proveriti samo naziv
	if err != nil {
		return nil, fmt.Errorf("failed to query all draft rights: %v", err)
	}
	defer rows.Close()

	var draftRights []model.DraftRight
	for rows.Next() {
		var draftRight model.DraftRight
		if err := rows.Scan(&draftRight.IdPrava, &draftRight.ImeIgrPrava, &draftRight.PrezimeIgrPrava,
			&draftRight.PozicijaIgrPrava, &draftRight.StatusPrava); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
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
	row := repo.db.QueryRow("SELECT * FROM PRAVA_NA_IGRACA WHERE IDPRAVA = :1", id) // Proveriti samo naziv
	if err := row.Scan(&draftRight.IdPrava, &draftRight.ImeIgrPrava, &draftRight.PrezimeIgrPrava,
		&draftRight.PozicijaIgrPrava, &draftRight.StatusPrava); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &draftRight, nil
}
