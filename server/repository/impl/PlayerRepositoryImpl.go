package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type playerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) repository.PlayerRepository {
	return &playerRepository{db}
}

func (repo *playerRepository) GetAll() ([]model.Player, error) {
	rows, err := repo.db.Query("SELECT * FROM IGRAC")
	if err != nil {
		return nil, fmt.Errorf("failed to query all players: %v", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var player model.Player
		if err := rows.Scan(&player.Id, &player.Ime, &player.Prezime, &player.Email, &player.DatRodj,
			&player.Lozinka, &player.Uloga, &player.VisIgr, &player.TezIgr, &player.PozIgr); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return players, nil
}

func (repo *playerRepository) GetByID(id int) (*model.Player, error) {
	var player model.Player
	row := repo.db.QueryRow("SELECT * FROM PIK WHERE IDPIK = :1", id)
	if err := row.Scan(&player.Id, &player.Ime, &player.Prezime, &player.Email, &player.DatRodj,
		&player.Lozinka, &player.Uloga, &player.VisIgr, &player.TezIgr, &player.PozIgr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &player, nil
}

func (repo *playerRepository) GetAllByTeamID(teamId int) ([]model.Player, error) {
	rows, err := repo.db.Query("SELECT * FROM Player WHERE IDTIM = :1", teamId) // PROVERITI DA LI JE IDTIM DOBRO
	if err != nil {
		return nil, fmt.Errorf("failed to query all players by team id: %v", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var player model.Player
		if err := rows.Scan(&player.Id, &player.Ime, &player.Prezime, &player.Email, &player.DatRodj,
			&player.Lozinka, &player.Uloga, &player.VisIgr, &player.TezIgr, &player.PozIgr); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return players, nil
}
