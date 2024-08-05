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
	rows, err := repo.db.Query(`SELECT IDPRAVA, IDTIM, IDREGRUT, IDPIK, NEDODLISTPRAVA, TRGLISTPRAVA
									  FROM PRAVANAIGRACA`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all draft rights: %v", err)
	}
	defer rows.Close()

	var draftRights []model.DraftRight
	for rows.Next() {
		var draftRightDAO model.DraftRightDAO
		if err := rows.Scan(&draftRightDAO.IdPrava, &draftRightDAO.IdTim, &draftRightDAO.IdRegrut, &draftRightDAO.IdPik,
			&draftRightDAO.NedodListPrava, &draftRightDAO.TrgListPrava); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		draftRight := &model.DraftRight{}
		draftRight.FromDAO(&draftRightDAO)

		draftRights = append(draftRights, *draftRight)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return draftRights, nil
}

func (repo *draftRightRepository) GetByID(id int) (*model.DraftRight, error) {
	var draftRightDAO model.DraftRightDAO
	row := repo.db.QueryRow(`SELECT IDPRAVA, IDTIM, IDREGRUT, IDPIK, NEDODLISTPRAVA, TRGLISTPRAVA
									FROM PRAVANAIGRACA
									WHERE IDPRAVA = :1`, id)
	if err := row.Scan(&draftRightDAO.IdPrava, &draftRightDAO.IdTim, &draftRightDAO.IdRegrut, &draftRightDAO.IdPik,
		&draftRightDAO.NedodListPrava, &draftRightDAO.TrgListPrava); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	draftRight := &model.DraftRight{}
	draftRight.FromDAO(&draftRightDAO)

	return draftRight, nil
}

func (repo *draftRightRepository) GetAllByTeamID(teamID int) ([]model.DraftRight, error) {
	rows, err := repo.db.Query(`SELECT IDPRAVA, IDTIM, IDREGRUT, IDPIK, NEDODLISTPRAVA, TRGLISTPRAVA
									  FROM PRAVANAIGRACA
									  WHERE IDTIM = :1`, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var draftRights []model.DraftRight
	for rows.Next() {
		var draftRightDAO model.DraftRightDAO
		if err := rows.Scan(&draftRightDAO.IdPrava, &draftRightDAO.IdTim, &draftRightDAO.IdRegrut, &draftRightDAO.IdPik,
			&draftRightDAO.NedodListPrava, &draftRightDAO.TrgListPrava); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		draftRight := &model.DraftRight{}
		draftRight.FromDAO(&draftRightDAO)

		draftRights = append(draftRights, *draftRight)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return draftRights, nil
}

func (repo *draftRightRepository) GetAllAvailableByTeamID(teamID int) ([]model.DraftRight, error) {
	rows, err := repo.db.Query(`SELECT IDPRAVA, IDTIM, IDREGRUT, IDPIK, NEDODLISTPRAVA, TRGLISTPRAVA
									  FROM PRAVANAIGRACA
									  WHERE IDTIM = :1 AND NEDODLISTPRAVA = 'FALSE'`, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all picks by team id: %v", err)
	}
	defer rows.Close()

	var draftRights []model.DraftRight
	for rows.Next() {
		var draftRightDAO model.DraftRightDAO
		if err := rows.Scan(&draftRightDAO.IdPrava, &draftRightDAO.IdTim, &draftRightDAO.IdRegrut, &draftRightDAO.IdPik,
			&draftRightDAO.NedodListPrava, &draftRightDAO.TrgListPrava); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		draftRight := &model.DraftRight{}
		draftRight.FromDAO(&draftRightDAO)

		draftRights = append(draftRights, *draftRight)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return draftRights, nil
}

func (repo *draftRightRepository) Update(draftRights *model.DraftRight) error {
	untouchable, tradeable := fromDraftRightsLists(draftRights)
	_, err := repo.db.Exec("UPDATE PravaNaIgraca SET IDTIM = :1, NEDODLISTPRAVA = :2, TRGLISTPRAVA = :3 WHERE IDPRAVA = :4",
		draftRights.TeamId, untouchable, tradeable, draftRights.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update draft rights: %v", err)
	}
	return nil
}

func (repo *draftRightRepository) AddToWishlist(draftRight *model.DraftRight, teamId int) error {
	_, err := repo.db.Exec(`INSERT INTO ZELJATIMA VALUES (0, SYSDATE, 'IGRAC', 2, :1, NULL, NULL, :2)`, // 2 is Draft Right type
		draftRight.ID, teamId)
	if err != nil {
		return fmt.Errorf("failed to add draft rights to the wishlist: %v", err)
	}
	return nil
}

func (repo *draftRightRepository) RemoveFromWishlist(draftRight *model.DraftRight, teamId int) error {
	_, err := repo.db.Exec(`DELETE FROM ZELJATIMA WHERE IDPRAVA = :1 AND IDTIM = :2`, draftRight.ID, teamId)
	if err != nil {
		return fmt.Errorf("failed to remove draft rights from the wishlist: %v", err)
	}
	return nil
}

func fromDraftRightsLists(draftRights *model.DraftRight) (string, string) {
	var untouchable, tradeable string
	if draftRights.Untouchable {
		untouchable = "TRUE"
	} else {
		untouchable = "FALSE"
	}
	if draftRights.Tradeable {
		tradeable = "TRUE"
	} else {
		tradeable = "FALSE"
	}
	return untouchable, tradeable
}
