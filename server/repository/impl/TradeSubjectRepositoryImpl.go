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
		var tradeSubject model.TradeSubject
		var tradeType string
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&tradeSubject.IdPredTrg, &tradeType, &tradeSubject.IdPrava, &tradeSubject.IdIgrac,
			&tradeSubject.IdZahTrg, &tradeSubject.IdPik); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		mapTradeSubjectEnumsForReading(tradeType, &tradeSubject)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubject)

		tradeSubjects = append(tradeSubjects, tradeSubject)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) GetByID(id int) (*model.TradeSubject, error) {
	var tradeSubject model.TradeSubject
	var tradeType string
	var idPrava, idIgrac, idPik sql.NullInt64
	row := repo.db.QueryRow("SELECT * FROM PredmetTrgovine WHERE IDPREDTRG = :1", id)
	if err := row.Scan(&tradeSubject.IdPredTrg, &tradeType, &tradeSubject.IdPrava, &tradeSubject.IdIgrac,
		&tradeSubject.IdZahTrg, &tradeSubject.IdPik); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	mapTradeSubjectEnumsForReading(tradeType, &tradeSubject)
	mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubject)

	return &tradeSubject, nil
}

func (repo *tradeSubjectRepository) GetAllByTradeProposalID(tradeProposalID int) ([]model.TradeSubject, error) {
	rows, err := repo.db.Query("SELECT * FROM PredmetTrgovine WHERE IDZAHTRG = :1", tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeSubjects []model.TradeSubject
	for rows.Next() {
		var tradeSubject model.TradeSubject
		var tradeType string
		var idPrava, idIgrac, idPik sql.NullInt64
		if err := rows.Scan(&tradeSubject.IdPredTrg, &tradeType, &idPrava, &idIgrac, &tradeSubject.IdZahTrg, &idPik); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		mapTradeSubjectEnumsForReading(tradeType, &tradeSubject)
		mapNullableAttributes(idPrava, idIgrac, idPik, &tradeSubject)

		tradeSubjects = append(tradeSubjects, tradeSubject)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeSubjects, nil
}

func (repo *tradeSubjectRepository) Create(tradeSubject *model.TradeSubject) error {
	if tradeSubject.TipPredTrg == 0 {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.IdPredTrg, "Igrac", nil,
			tradeSubject.IdIgrac, tradeSubject.IdZahTrg, nil)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	} else if tradeSubject.TipPredTrg == 1 {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.IdPredTrg, "Pik", nil, nil, tradeSubject.IdZahTrg, tradeSubject.IdPik)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	} else {
		_, err := repo.db.Exec("INSERT INTO PredmetTrgovine (IDPREDTRG, TIPPREDTRG, IDPRAVA, IDIGRAC, IDZAHTRG, "+
			"IDPIK) VALUES (:1, :2, :3, :4, :5, :6)", tradeSubject.IdPredTrg, "PravaNaIgraca", tradeSubject.IdPrava, nil, tradeSubject.IdZahTrg, nil)
		if err != nil {
			return fmt.Errorf("failed to create a trade proposal: %v", err)
		}
		return nil
	}
}

func mapTradeSubjectEnumsForReading(tradeType string, tradeSubject *model.TradeSubject) {
	switch tradeType {
	case "Igrac":
		tradeSubject.TipPredTrg = 0
	case "Pik":
		tradeSubject.TipPredTrg = 1
	default:
		tradeSubject.TipPredTrg = 2
	}
}

func mapTradeSubjectEnumsForWriting(tradeSubject *model.TradeSubject) string {
	switch tradeSubject.TipPredTrg {
	case 0:
		return "Igrac"
	case 1:
		return "Pik"
	default:
		return "PravaNaIgraca"
	}
}

func mapNullableAttributes(idPrava sql.NullInt64, idIgrac sql.NullInt64, idPik sql.NullInt64, tradeSubject *model.TradeSubject) {
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
