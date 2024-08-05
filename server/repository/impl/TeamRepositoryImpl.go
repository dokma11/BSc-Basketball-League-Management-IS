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
		var teamDAO model.TeamDAO
		if err := rows.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		team := &model.Team{}
		team.FromDAO(&teamDAO)

		teams = append(teams, *team)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return teams, nil
}

func (repo *teamRepository) GetByID(id int) (*model.Team, error) {
	var teamDAO model.TeamDAO
	row := repo.db.QueryRow("SELECT * FROM TIM WHERE IDTIM = :1", id)
	if err := row.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	team := &model.Team{}
	team.FromDAO(&teamDAO)

	return team, nil
}

func (repo *teamRepository) GetByUserID(userID int) (*model.Team, error) {
	var teamDAO model.TeamDAO
	row := repo.db.QueryRow(`SELECT T.IDTIM, T.NAZTIM, T.GODOSNTIM, T.LOKTIM
								   FROM TIM T, UGOVOR U, ZAPOSLENI Z
								   WHERE Z.IDUGO = U.IDUGO AND U.IDTIM = T.IDTIM AND Z.ID = :1`, userID)
	if err := row.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	team := &model.Team{}
	team.FromDAO(&teamDAO)

	return team, nil
}

func (repo *teamRepository) GetPlayerTradeDestination(tradeSubjectID int) (*model.Team, error) {
	var teamDAO model.TeamDAO
	row := repo.db.QueryRow(`WITH TIM_PREDMETA AS (SELECT U.IDTIM
								   FROM ZAPOSLENI ZAP, UGOVOR U, PREDMETTRGOVINE P
								   WHERE P.IDPREDTRG = :1 AND ZAP.ID = P.IDIGRAC AND ZAP.IDUGO = U.IDUGO) 
								   SELECT T.IDTIM, T.NAZTIM, T.GODOSNTIM, T.LOKTIM
								   FROM ZAHTEVZATRGOVINU Z, ZAPOSLENI ZAP_MEN, PREDMETTRGOVINE P, UGOVOR U, TIM T, TIM_PREDMETA TP
								   WHERE P.IDPREDTRG = :2 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPRIM = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != TP.IDTIM AND U.IDTIM = T.IDTIM OR 
								   P.IDPREDTRG = :3 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPOS = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != TP.IDTIM AND U.IDTIM = T.IDTIM`, tradeSubjectID, tradeSubjectID, tradeSubjectID)
	if err := row.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println(err)
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	team := &model.Team{}
	team.FromDAO(&teamDAO)

	return team, nil
}

func (repo *teamRepository) GetPickTradeDestination(tradeSubjectID int) (*model.Team, error) {
	var teamDAO model.TeamDAO
	row := repo.db.QueryRow(`WITH PIKOV_TIM AS (SELECT T.IDTIM 
								   FROM TIM T, PIK P, PREDMETTRGOVINE PT
								   WHERE PT.IDPREDTRG = :1 AND PT.IDPIK = P.IDPIK AND T.IDTIM = P.IDTIM)
								   SELECT T.IDTIM, T.NAZTIM, T.GODOSNTIM, T.LOKTIM
								   FROM ZAHTEVZATRGOVINU Z, ZAPOSLENI ZAP_MEN, PREDMETTRGOVINE P, UGOVOR U, TIM T, PIKOV_TIM PT
								   WHERE P.IDPREDTRG = :2 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPRIM = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != PT.IDTIM AND U.IDTIM = T.IDTIM OR 
								   P.IDPREDTRG = :3 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPOS = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != PT.IDTIM AND U.IDTIM = T.IDTIM`, tradeSubjectID, tradeSubjectID, tradeSubjectID)
	if err := row.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	team := &model.Team{}
	team.FromDAO(&teamDAO)

	return team, nil
}

func (repo *teamRepository) GetDraftRightsTradeDestination(tradeSubjectID int) (*model.Team, error) {
	var teamDAO model.TeamDAO
	row := repo.db.QueryRow(`WITH PRAVA_OD_TIMA AS (SELECT T.IDTIM 
								   FROM TIM T, PRAVANAIGRACA P, PREDMETTRGOVINE PT
								   WHERE PT.IDPREDTRG = :1 AND PT.IDPRAVA = P.IDPRAVA AND T.IDTIM = P.IDTIM)
								   SELECT T.IDTIM, T.NAZTIM, T.GODOSNTIM, T.LOKTIM
								   FROM ZAHTEVZATRGOVINU Z, ZAPOSLENI ZAP_MEN, PREDMETTRGOVINE P, UGOVOR U, TIM T, PRAVA_OD_TIMA POT
								   WHERE P.IDPREDTRG = :2 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPRIM = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != POT.IDTIM AND U.IDTIM = T.IDTIM OR 
								   P.IDPREDTRG = :3 AND P.IDZAHTRG = Z.IDZAHTRG AND Z.IDMENADZERPOS = ZAP_MEN.ID AND ZAP_MEN.IDUGO = U.IDUGO AND U.IDTIM != POT.IDTIM AND U.IDTIM = T.IDTIM`, tradeSubjectID, tradeSubjectID, tradeSubjectID)
	if err := row.Scan(&teamDAO.IdTim, &teamDAO.NazTim, &teamDAO.GodOsnTim, &teamDAO.LokTim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	team := &model.Team{}
	team.FromDAO(&teamDAO)

	return team, nil
}

func (repo *teamRepository) GetWishlistByTeamID(teamID int) ([]model.WishlistAsset, error) {
	rows, err := repo.db.Query(`SELECT * FROM ZELJATIMA WHERE IDTIM = :1`, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to query team's wishlist: %v", err)
	}
	defer rows.Close()

	var wishlist []model.WishlistAsset
	for rows.Next() {
		var wish model.WishlistAsset
		var draftRightsID, pickID, playerID sql.NullInt64
		if err := rows.Scan(&wish.IdZeljTim, &wish.DatDodZeljTim, &wish.BelesZeljTim, &wish.IdTipZelje,
			&draftRightsID, &pickID, &playerID, &wish.IdTim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromNullables(draftRightsID, pickID, playerID, &wish)

		wishlist = append(wishlist, wish)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return wishlist, nil
}

func fromNullables(draftRightsID sql.NullInt64, pickID sql.NullInt64, playerID sql.NullInt64, wishlist *model.WishlistAsset) {
	if draftRightsID.Valid {
		wishlist.IdPrava = draftRightsID.Int64
	} else if pickID.Valid {
		wishlist.IdPik = pickID.Int64
	} else if playerID.Valid {
		wishlist.IdIgrac = playerID.Int64
	}
}
