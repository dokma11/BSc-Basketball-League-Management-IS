package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type tradeSubjectRepository struct {
	db *sql.DB
}

func NewTradeSubjectRepository(db *sql.DB) repository.TradeSubjectRepository {
	return &tradeSubjectRepository{db}
}

func (repo *tradeSubjectRepository) GetAll() ([]model.TradeSubject, error) {
	rows, err := repo.db.Query("SELECT * FROM PredmetTrgovine")
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade subjects: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubject
	for rows.Next() {
		var tradeSubjectDAO model.TradeSubjectDAO
		var tradeType string
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&tradeSubjectDAO.IdPredTrg, &tradeType, &tradeSubjectDAO.IdPrava, &tradeSubjectDAO.IdIgrac,
			&tradeSubjectDAO.IdZahTrg, &tradeSubjectDAO.IdPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
		tradeSubject := &model.TradeSubject{}
		tradeSubject.FromDAO(&tradeSubjectDAO)

		tradeSubjects = append(tradeSubjects, *tradeSubject)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) GetByID(id int) (*model.TradeSubject, error) {
	var tradeSubjectDAO model.TradeSubjectDAO
	var tradeType string
	var idPrava, idIgrac, idPik sql.NullInt64
	row := repo.db.QueryRow("SELECT * FROM PredmetTrgovine WHERE IDPREDTRG = :1", id)
	if err := row.Scan(&tradeSubjectDAO.IdPredTrg, &tradeType, &tradeSubjectDAO.IdPrava, &tradeSubjectDAO.IdIgrac,
		&tradeSubjectDAO.IdZahTrg, &tradeSubjectDAO.IdPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
	mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
	tradeSubject := &model.TradeSubject{}
	tradeSubject.FromDAO(&tradeSubjectDAO)

	return tradeSubject, nil
}

func (repo *tradeSubjectRepository) GetAllByTradeProposalID(tradeProposalID int) ([]model.TradeSubject, error) {
	rows, err := repo.db.Query("SELECT * FROM PredmetTrgovine WHERE IDZAHTRG = :1", tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubject
	for rows.Next() {
		var tradeSubjectDAO model.TradeSubjectDAO
		var tradeType string
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&tradeSubjectDAO.IdPredTrg, &tradeType, &idPrava, &idIgrac, &tradeSubjectDAO.IdZahTrg, &idPik); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
		tradeSubject := &model.TradeSubject{}
		tradeSubject.FromDAO(&tradeSubjectDAO)

		tradeSubjects = append(tradeSubjects, *tradeSubject)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) GetPlayerTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error) {
	rows, err := repo.db.Query(`SELECT U.IDTIM, P.IDPREDTRG, P.TIPPREDTRG, P.IDPRAVA, P.IDIGRAC, P.IDZAHTRG, P.IDPIK
									  FROM ZAPOSLENI ZAP, UGOVOR U, PREDMETTRGOVINE P, ZAHTEVZATRGOVINU Z
									  WHERE ZAP.ID = P.IDIGRAC AND ZAP.IDUGO = U.IDUGO AND Z.IDZAHTRG = P.IDZAHTRG AND Z.IDZAHTRG = :1`, tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubjectDetailsResponseDTO
	for rows.Next() {
		var tradeSubjectDAO model.TradeSubjectDAO
		var tradeType string
		var teamId int64
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&teamId, &tradeSubjectDAO.IdPredTrg, &tradeType, &idPrava, &idIgrac, &tradeSubjectDAO.IdZahTrg, &idPik); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
		tradeSubject := &model.TradeSubject{}
		tradeSubject.FromDAO(&tradeSubjectDAO)

		var tradeSubjectDTO model.TradeSubjectDetailsResponseDTO
		mapToTradeSubjectDTO(*tradeSubject, &tradeSubjectDTO, teamId)

		tradeSubjects = append(tradeSubjects, tradeSubjectDTO)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) GetPickTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error) {
	rows, err := repo.db.Query(`SELECT PI.IDTIM, P.IDPREDTRG, P.TIPPREDTRG, P.IDPRAVA, P.IDIGRAC, P.IDZAHTRG, P.IDPIK
									  FROM PREDMETTRGOVINE P, ZAHTEVZATRGOVINU Z, PIK PI
									  WHERE Z.IDZAHTRG = P.IDZAHTRG AND PI.IDPIK = P.IDPIK AND Z.IDZAHTRG = :1`, tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubjectDetailsResponseDTO
	for rows.Next() {
		var tradeSubjectDAO model.TradeSubjectDAO
		var tradeType string
		var teamId int64
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&teamId, &tradeSubjectDAO.IdPredTrg, &tradeType, &idPrava, &idIgrac, &tradeSubjectDAO.IdZahTrg, &idPik); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
		tradeSubject := &model.TradeSubject{}
		tradeSubject.FromDAO(&tradeSubjectDAO)

		var tradeSubjectDTO model.TradeSubjectDetailsResponseDTO
		mapToTradeSubjectDTO(*tradeSubject, &tradeSubjectDTO, teamId)

		tradeSubjects = append(tradeSubjects, tradeSubjectDTO)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) GetDraftRightsTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error) {
	rows, err := repo.db.Query(`SELECT PI.IDTIM, P.IDPREDTRG, P.TIPPREDTRG, P.IDPRAVA, P.IDIGRAC, P.IDZAHTRG, P.IDPIK
									  FROM PREDMETTRGOVINE P, ZAHTEVZATRGOVINU Z, PRAVANAIGRACA PI
									  WHERE Z.IDZAHTRG = P.IDZAHTRG AND PI.IDPRAVA = P.IDPRAVA AND Z.IDZAHTRG = :1`, tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubjectDetailsResponseDTO
	for rows.Next() {
		var tradeSubjectDAO model.TradeSubjectDAO
		var tradeType string
		var teamId int64
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&teamId, &tradeSubjectDAO.IdPredTrg, &tradeType, &idPrava, &idIgrac, &tradeSubjectDAO.IdZahTrg, &idPik); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTradeSubjectTypeForReading(tradeType, &tradeSubjectDAO)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubjectDAO)
		tradeSubject := &model.TradeSubject{}
		tradeSubject.FromDAO(&tradeSubjectDAO)

		var tradeSubjectDTO model.TradeSubjectDetailsResponseDTO
		mapToTradeSubjectDTO(*tradeSubject, &tradeSubjectDTO, teamId)

		tradeSubjects = append(tradeSubjects, tradeSubjectDTO)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) Create(tradeSubject *model.TradeSubject) error {
	if tradeSubject.Type == 0 {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.ID, "Igrac", nil,
			tradeSubject.PlayerId, tradeSubject.TradeProposalId, nil)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	} else if tradeSubject.Type == 1 {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.ID, "Pik", nil, nil, tradeSubject.TradeProposalId, tradeSubject.PickId)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	} else {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.ID, "PravaNaIgraca", tradeSubject.DraftRightsId, nil, tradeSubject.TradeProposalId, nil)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	}
}

func fromTradeSubjectTypeForReading(tradeType string, tradeSubject *model.TradeSubjectDAO) {
	switch tradeType {
	case "Igrac":
		tradeSubject.TipPredTrg = 0
	case "Pik":
		tradeSubject.TipPredTrg = 1
	default:
		tradeSubject.TipPredTrg = 2
	}
}

func fromTradeSubjectTypeForWriting(tradeSubject *model.TradeSubject) string {
	switch tradeSubject.Type {
	case 0:
		return "Igrac"
	case 1:
		return "Pik"
	default:
		return "PravaNaIgraca"
	}
}

func mapNullableAttributes(idPrava sql.NullInt64, idIgrac sql.NullInt64, idPik sql.NullInt64, tradeSubject *model.TradeSubjectDAO) {
	if idPrava.Valid {
		tradeSubject.IdPrava = idPrava.Int64
	}
	if idIgrac.Valid {
		tradeSubject.IdIgrac = idIgrac.Int64
	}
	if idPik.Valid {
		tradeSubject.IdPik = idPik.Int64
	}
}

func mapToTradeSubjectDTO(tradeSubject model.TradeSubject, tradeSubjectDTO *model.TradeSubjectDetailsResponseDTO, teamId int64) {
	tradeSubjectDTO.IdPredTrg = tradeSubject.ID
	tradeSubjectDTO.TipPredTrg = tradeSubject.Type
	tradeSubjectDTO.IdPrava = tradeSubject.DraftRightsId
	tradeSubjectDTO.IdTim = teamId
	tradeSubjectDTO.IdPik = tradeSubject.PickId
	tradeSubjectDTO.IdZahTrg = tradeSubject.TradeProposalId
	tradeSubjectDTO.IdIgrac = tradeSubject.PlayerId
}
