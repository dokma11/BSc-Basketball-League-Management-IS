package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type tradeProposalRepository struct {
	db *sql.DB
}

func NewTradeProposalRepository(db *sql.DB) repository.TradeProposalRepository {
	return &tradeProposalRepository{db}
}

func (repo *tradeProposalRepository) GetAll() ([]model.TradeProposal, error) {
	rows, err := repo.db.Query("SELECT * FROM ZahtevZaTrgovinu") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeProposals []model.TradeProposal
	for rows.Next() {
		var tradeProposal model.TradeProposal
		var tradeType string
		var status string
		if err := rows.Scan(&tradeProposal.IdZahTrg, &tradeProposal.DatZahTrg, &tradeType, &status, &tradeProposal.RazlogOdbij,
			&tradeProposal.IdMenadzerPos, &tradeProposal.IdMenadzerPrim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if tradeType == "PLAYER_PLAYER" {
			tradeProposal.TipZahTrg = 0
		} else if tradeType == "PLAYER_PICK" {
			tradeProposal.TipZahTrg = 1
		} else if tradeType == "PICK_PICK" {
			tradeProposal.TipZahTrg = 2
		}

		if status == "IN_PROGRESS" {
			tradeProposal.StatusZahTrg = 0
		} else if status == "ACCEPTED" {
			tradeProposal.StatusZahTrg = 1
		} else if status == "DECLINED" {
			tradeProposal.StatusZahTrg = 2
		} else if status == "CANCELLED" {
			tradeProposal.StatusZahTrg = 3
		}

		tradeProposals = append(tradeProposals, tradeProposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeProposals, nil
}

func (repo *tradeProposalRepository) GetByID(id int) (*model.TradeProposal, error) {
	var tradeProposal model.TradeProposal
	var tradeType string
	var status string
	row := repo.db.QueryRow("SELECT * FROM ZahtevZaTrgovinu WHERE IDZAHTRG = :1", id)
	if err := row.Scan(&tradeProposal.IdZahTrg, &tradeProposal.DatZahTrg, &tradeType, &status, &tradeProposal.RazlogOdbij,
		&tradeProposal.IdMenadzerPos, &tradeProposal.IdMenadzerPrim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	if tradeType == "PLAYER_PLAYER" {
		tradeProposal.TipZahTrg = 0
	} else if tradeType == "PLAYER_PICK" {
		tradeProposal.TipZahTrg = 1
	} else if tradeType == "PICK_PICK" {
		tradeProposal.TipZahTrg = 2
	}

	if status == "IN_PROGRESS" {
		tradeProposal.StatusZahTrg = 0
	} else if status == "ACCEPTED" {
		tradeProposal.StatusZahTrg = 1
	} else if status == "DECLINED" {
		tradeProposal.StatusZahTrg = 2
	} else if status == "CANCELLED" {
		tradeProposal.StatusZahTrg = 3
	}

	return &tradeProposal, nil
}

func (repo *tradeProposalRepository) GetAllByTeamID(teamID int) ([]model.TradeProposal, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM ZahtevZaTrgovinu") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeProposals []model.TradeProposal
	for rows.Next() {
		var tradeProposal model.TradeProposal
		var tradeType string
		var status string
		if err := rows.Scan(&tradeProposal.IdZahTrg, &tradeProposal.DatZahTrg, &tradeType, &status, &tradeProposal.RazlogOdbij,
			&tradeProposal.IdMenadzerPos, &tradeProposal.IdMenadzerPrim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if tradeType == "PLAYER_PLAYER" {
			tradeProposal.TipZahTrg = 0
		} else if tradeType == "PLAYER_PICK" {
			tradeProposal.TipZahTrg = 1
		} else if tradeType == "PICK_PICK" {
			tradeProposal.TipZahTrg = 2
		}

		if status == "IN_PROGRESS" {
			tradeProposal.StatusZahTrg = 0
		} else if status == "ACCEPTED" {
			tradeProposal.StatusZahTrg = 1
		} else if status == "DECLINED" {
			tradeProposal.StatusZahTrg = 2
		} else if status == "CANCELLED" {
			tradeProposal.StatusZahTrg = 3
		}

		tradeProposals = append(tradeProposals, tradeProposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeProposals, nil
}

func (repo *tradeProposalRepository) Create(tradeProposal *model.TradeProposal) error {
	_, err := repo.db.Exec("INSERT INTO ZahtevZaTrgovinu (IDZAHTRG, DATZAHTRG, TIPZAHTRG, STATUSZAHTRG, RAZLOGODBIJ, "+
		"IDMENADZERPOS, IDMENADZERPRIM) "+
		"VALUES (:1, :2, :3, :4, :5, :6, :7)", tradeProposal.IdZahTrg, tradeProposal.DatZahTrg, tradeProposal.TipZahTrg,
		tradeProposal.StatusZahTrg, tradeProposal.RazlogOdbij, tradeProposal.IdMenadzerPos, tradeProposal.IdMenadzerPrim)
	if err != nil {
		return fmt.Errorf("failed to create a trade proposal: %v", err)
	}
	return nil
}

func (repo *tradeProposalRepository) Update(tradeProposal *model.TradeProposal) error {
	_, err := repo.db.Exec("UPDATE ZahtevZaTrgovinu SET DATZAHTRG = :1, TIPZAHTRG = :2, STATUSZAHTRG = :3"+
		", RAZLOGODBIJ = :4 WHERE IDZAHTRG = :5", tradeProposal.DatZahTrg, tradeProposal.TipZahTrg,
		tradeProposal.StatusZahTrg, tradeProposal.RazlogOdbij, tradeProposal.IdZahTrg)
	if err != nil {
		return fmt.Errorf("failed to update tim: %v", err)
	}
	return nil
}
