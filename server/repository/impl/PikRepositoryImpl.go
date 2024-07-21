package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type pikRepository struct {
	db *sql.DB
}

func NewPikRepository(db *sql.DB) repository.PikRepository {
	return &pikRepository{db}
}

func (repo *pikRepository) GetAll() ([]model.Pik, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK")
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks: %v", err)
	}
	defer rows.Close()

	var picks []model.Pik
	for rows.Next() {
		var pick model.Pik
		if err := rows.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pikRepository) GetByID(id int) (*model.Pik, error) {
	var pick model.Pik
	row := repo.db.QueryRow("SELECT * FROM PIK WHERE IDPIK = :1", id)
	if err := row.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &pick, nil
}

func (repo *pikRepository) GetAllByTeamID(teamId int) ([]model.Pik, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE IDTIM = :1", teamId)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var picks []model.Pik
	for rows.Next() {
		var pick model.Pik
		if err := rows.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pikRepository) GetAllByYear(year string) ([]model.Pik, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE GODPIK = :1", year)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by year: %v", err)
	}
	defer rows.Close()

	var picks []model.Pik
	for rows.Next() {
		var pick model.Pik
		if err := rows.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}
