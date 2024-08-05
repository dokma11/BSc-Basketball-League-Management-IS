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
		var pickDAO model.PickDAO
		var managerID sql.NullInt64
		if err := rows.Scan(&pickDAO.IdPik, &pickDAO.RedBrPik, &pickDAO.BrRunPik, &pickDAO.GodPik, &managerID,
			&pickDAO.IdTim, &pickDAO.NedodListPik, &pickDAO.TrgListPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if managerID.Valid {
			pickDAO.IdMenadzer = managerID.Int64
		}

		pick := &model.Pick{}
		pick.FromDAO(&pickDAO)

		picks = append(picks, *pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pickRepository) GetByID(id int) (*model.Pick, error) {
	var pickDAO model.PickDAO
	var managerID sql.NullInt64
	row := repo.db.QueryRow("SELECT * FROM PIK WHERE IDPIK = :1", id)
	if err := row.Scan(&pickDAO.IdPik, &pickDAO.RedBrPik, &pickDAO.BrRunPik, &pickDAO.GodPik, &managerID,
		&pickDAO.IdTim, &pickDAO.NedodListPik, &pickDAO.TrgListPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	if managerID.Valid {
		pickDAO.IdMenadzer = managerID.Int64
	}

	pick := &model.Pick{}
	pick.FromDAO(&pickDAO)

	return pick, nil
}

// TODO: FIltrirati pikove da ne bduud oni untouchable
func (repo *pickRepository) GetAllByTeamID(teamId int) ([]model.Pick, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE IDTIM = :1 AND IDMENADZER IS NULL", teamId) // IDMENADZER IS NULL means that the pick has not been used
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var picks []model.Pick
	for rows.Next() {
		var pickDAO model.PickDAO
		var managerID sql.NullInt64
		if err := rows.Scan(&pickDAO.IdPik, &pickDAO.RedBrPik, &pickDAO.BrRunPik, &pickDAO.GodPik, &managerID,
			&pickDAO.IdTim, &pickDAO.NedodListPik, &pickDAO.TrgListPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if managerID.Valid {
			pickDAO.IdMenadzer = managerID.Int64
		}

		pick := &model.Pick{}
		pick.FromDAO(&pickDAO)

		picks = append(picks, *pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pickRepository) GetAllAvailableByTeamID(teamId int) ([]model.Pick, error) {
	rows, err := repo.db.Query("SELECT * FROM PIK WHERE IDTIM = :1 AND IDMENADZER IS NULL AND NEDODLISTPIK = 'FALSE'", teamId)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var picks []model.Pick
	for rows.Next() {
		var pickDAO model.PickDAO
		var managerID sql.NullInt64
		if err := rows.Scan(&pickDAO.IdPik, &pickDAO.RedBrPik, &pickDAO.BrRunPik, &pickDAO.GodPik, &managerID,
			&pickDAO.IdTim, &pickDAO.NedodListPik, &pickDAO.TrgListPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if managerID.Valid {
			pickDAO.IdMenadzer = managerID.Int64
		}

		pick := &model.Pick{}
		pick.FromDAO(&pickDAO)

		picks = append(picks, *pick)
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
		var pickDAO model.PickDAO
		var managerID sql.NullInt64
		if err := rows.Scan(&pickDAO.IdPik, &pickDAO.RedBrPik, &pickDAO.BrRunPik, &pickDAO.GodPik, &managerID,
			&pickDAO.IdTim, &pickDAO.NedodListPik, &pickDAO.TrgListPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if managerID.Valid {
			pickDAO.IdMenadzer = managerID.Int64
		}

		pick := &model.Pick{}
		pick.FromDAO(&pickDAO)

		picks = append(picks, *pick)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return picks, nil
}

func (repo *pickRepository) Update(pick *model.Pick) error {
	untouchable, tradeable := fromPickLists(pick)
	_, err := repo.db.Exec("UPDATE PIK SET IDTIM = :1, NEDODLISTPIK = :2, TRGLISTPIK = :3 WHERE IDPIK = :4",
		pick.TeamId, untouchable, tradeable, pick.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update pick: %v", err)
	}
	return nil
}

func fromPickLists(pick *model.Pick) (string, string) {
	var untouchable, tradeable string
	if pick.Untouchable {
		untouchable = "TRUE"
	} else if !pick.Untouchable {
		untouchable = "FALSE"
	}

	if pick.Tradeable {
		tradeable = "TRUE"
	} else if !pick.Tradeable {
		tradeable = "FALSE"
	}

	return untouchable, tradeable
}
