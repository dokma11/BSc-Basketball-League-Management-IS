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
	rows, err := repo.db.Query(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, I.VISIGR, I.TEZIGR, I.POZIGR, I.NEDODLISTIGR, I.TRGLISTIGR
								   	  FROM IGRAC I, KORISNIK K
								      WHERE I.ID = K.ID`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all players: %v", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var playerDAO model.PlayerDAO
		var role, position string
		if err := rows.Scan(&playerDAO.ID, &playerDAO.FirstName, &playerDAO.LastName, &playerDAO.Email, &playerDAO.DateOfBirth,
			&playerDAO.Password, &role, &playerDAO.VisIgr, &playerDAO.TezIgr, &position, &playerDAO.NedodListIgr, &playerDAO.TrgListIgr); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRoleAndPositionString(role, position, &playerDAO)
		player := &model.Player{}
		player.FromDAO(&playerDAO)

		players = append(players, *player)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return players, nil
}

func (repo *playerRepository) GetByID(id int) (*model.Player, error) {
	var playerDAO model.PlayerDAO
	var role, position string
	row := repo.db.QueryRow(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, I.VISIGR, I.TEZIGR, I.POZIGR, I.NEDODLISTIGR, I.TRGLISTIGR
								   FROM IGRAC I, KORISNIK K
								   WHERE I.ID = K.ID AND K.ID = :1`, id)
	if err := row.Scan(&playerDAO.ID, &playerDAO.FirstName, &playerDAO.LastName, &playerDAO.Email, &playerDAO.DateOfBirth,
		&playerDAO.Password, &role, &playerDAO.VisIgr, &playerDAO.TezIgr, &position, &playerDAO.NedodListIgr, &playerDAO.TrgListIgr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRoleAndPositionString(role, position, &playerDAO)
	player := &model.Player{}
	player.FromDAO(&playerDAO)

	return player, nil
}

func (repo *playerRepository) GetAllByTeamID(teamId int) ([]model.Player, error) {
	rows, err := repo.db.Query(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, I.VISIGR, I.TEZIGR, I.POZIGR, I.NEDODLISTIGR, I.TRGLISTIGR
									  FROM IGRAC I, ZAPOSLENI Z, KORISNIK K, UGOVOR U
									  WHERE I.ID = Z.ID AND I.ID = K.ID AND Z.IDUGO = U.IDUGO AND U.IDTIM = :1`, teamId)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to query all players by team id: %v", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var playerDAO model.PlayerDAO
		var role, position string
		if err := rows.Scan(&playerDAO.ID, &playerDAO.FirstName, &playerDAO.LastName, &playerDAO.Email, &playerDAO.DateOfBirth,
			&playerDAO.Password, &role, &playerDAO.VisIgr, &playerDAO.TezIgr, &position, &playerDAO.NedodListIgr, &playerDAO.TrgListIgr); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRoleAndPositionString(role, position, &playerDAO)
		player := &model.Player{}
		player.FromDAO(&playerDAO)

		players = append(players, *player)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return players, nil
}

func (repo *playerRepository) GetAllAvailableByTeamID(teamId int) ([]model.Player, error) {
	rows, err := repo.db.Query(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, I.VISIGR, I.TEZIGR, I.POZIGR, I.NEDODLISTIGR, I.TRGLISTIGR
									  FROM IGRAC I, ZAPOSLENI Z, KORISNIK K, UGOVOR U
									  WHERE I.ID = Z.ID AND I.ID = K.ID AND Z.IDUGO = U.IDUGO AND U.IDTIM = :1 AND I.NEDODLISTIGR = 'FALSE'`, teamId)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to query all players by team id: %v", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var playerDAO model.PlayerDAO
		var role, position string
		if err := rows.Scan(&playerDAO.ID, &playerDAO.FirstName, &playerDAO.LastName, &playerDAO.Email, &playerDAO.DateOfBirth,
			&playerDAO.Password, &role, &playerDAO.VisIgr, &playerDAO.TezIgr, &position, &playerDAO.NedodListIgr, &playerDAO.TrgListIgr); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRoleAndPositionString(role, position, &playerDAO)
		player := &model.Player{}
		player.FromDAO(&playerDAO)

		players = append(players, *player)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return players, nil
}

func (repo *playerRepository) Update(player *model.Player) error {
	untouchable, tradeable := fromLists(player)
	_, err := repo.db.Exec("UPDATE IGRAC SET NEDODLISTIGR = :1, TRGLISTIGR = :2 WHERE ID = :3",
		untouchable, tradeable, player.ID)
	if err != nil {
		return fmt.Errorf("failed to update trade proposal: %v", err)
	}
	return nil
}

func fromRoleAndPositionString(role string, position string, player *model.PlayerDAO) {
	if role == "Zaposleni" {
		player.Role = 1
	} else if role == "Regrut" {
		player.Role = 0
	}

	if position == "PG" {
		player.PozIgr = 0
	} else if position == "SG" {
		player.PozIgr = 1
	} else if position == "SF" {
		player.PozIgr = 2
	} else if position == "PF" {
		player.PozIgr = 3
	} else if position == "C" {
		player.PozIgr = 4
	}
}

func fromLists(player *model.Player) (string, string) {
	var untouchable, tradeable string
	if player.Untouchable {
		untouchable = "TRUE"
	} else if !player.Untouchable {
		untouchable = "FALSE"
	}

	if player.Tradeable {
		tradeable = "TRUE"
	} else if !player.Tradeable {
		tradeable = "FALSE"
	}

	return untouchable, tradeable
}
