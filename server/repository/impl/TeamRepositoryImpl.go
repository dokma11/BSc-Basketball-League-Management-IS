package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) repository.TeamRepository {
	return &teamRepository{db}
}

func (repo *teamRepository) GetAll() ([]model.Team, error) {
	rows, err := repo.db.Query("SELECT * FROM TIM")
	if err != nil {
		return nil, fmt.Errorf("failed to query all teams: %v", err)
	}
	defer rows.Close()

	var teams []model.Team
	for rows.Next() {
		var team model.Team
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

func (repo *teamRepository) GetByID(id int) (*model.Team, error) {
	var team model.Team
	row := repo.db.QueryRow("SELECT * FROM TIM WHERE IDTIM = :1", id)
	if err := row.Scan(&team.IdTim, &team.NazTim, &team.GodOsnTim, &team.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &team, nil
}

func (repo *teamRepository) GetByUserID(userID int) (*model.Team, error) {
	var team model.Team
	row := repo.db.QueryRow(`SELECT T.IDTIM, T.NAZTIM, T.GODOSNTIM, T.LOKTIM
								   FROM TIM T, UGOVOR U, ZAPOSLENI Z
								   WHERE Z.IDUGO = U.IDUGO AND U.IDTIM = T.IDTIM AND Z.ID = :1`, userID)
	if err := row.Scan(&team.IdTim, &team.NazTim, &team.GodOsnTim, &team.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &team, nil
}
