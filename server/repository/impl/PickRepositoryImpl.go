package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type pickRepository struct {
	db *sql.DB
}

func NewPickRepository(db *sql.DB) repository.PickRepository {
	return &pickRepository{db}
}

func (repo *pickRepository) GetAll() ([]model.Pick, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK")
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks: %v", err)
	}
	defer rows.Close()

	var picks []model.Pick
	for rows.Next() {
		var pick model.Pick
		if err := rows.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik, &pick.IdMenadzer, &pick.IdTim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pickRepository) GetByID(id int) (*model.Pick, error) {
	var pick model.Pick
	row := repo.db.QueryRow("SELECT * FROM PIK WHERE IDPIK = :1", id)
	if err := row.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik, &pick.IdMenadzer, &pick.IdTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &pick, nil
}

func (repo *pickRepository) GetAllByTeamID(teamId int) ([]model.Pick, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE IDTIM = :1", teamId)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var picks []model.Pick
	for rows.Next() {
		var pick model.Pick
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

func (repo *pickRepository) GetAllByYear(year string) ([]model.Pick, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE GODPIK = :1", year)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by year: %v", err)
	}
	defer rows.Close()

	var picks []model.Pick
	for rows.Next() {
		var pick model.Pick
		if err := rows.Scan(&pick.IdPik, &pick.RedBrPik, &pick.BrRunPik, &pick.GodPik, &pick.IdMenadzer, &pick.IdTim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		picks = append(picks, pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}
