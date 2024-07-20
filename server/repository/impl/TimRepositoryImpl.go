package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type timRepository struct {
	db *sql.DB
}

func NewTimRepository(db *sql.DB) repository.TimRepository {
	return &timRepository{db}
}

func (repo *timRepository) GetAll() ([]model.Tim, error) {
	rows, err := repo.db.Query("SELECT * FROM TIM")
	if err != nil {
		return nil, fmt.Errorf("failed to query all teams: %v", err)
	}
	defer rows.Close()

	var teams []model.Tim
	for rows.Next() {
		var team model.Tim
		if err := rows.Scan(&team.IdTim, &team.NazTim, &team.GodOsnTim, &team.LokTim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return teams, nil
}

func (repo *timRepository) GetByID(id int) (*model.Tim, error) {
	var team model.Tim
	row := repo.db.QueryRow("SELECT * FROM TIM WHERE IDTIM = :1", id)
	if err := row.Scan(&team.IdTim, &team.NazTim, &team.GodOsnTim, &team.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &team, nil
}
