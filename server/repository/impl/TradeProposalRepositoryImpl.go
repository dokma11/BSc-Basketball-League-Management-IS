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
	rows, err := repo.db.Query("SELECT * FROM ZahtevZaTrgovinu")
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeProposals []model.TradeProposal
	for rows.Next() {
		var tradeProposalDAO model.TradeProposalDAO
		var tradeType, status string
		if err := rows.Scan(&tradeProposalDAO.IdZahTrg, &tradeProposalDAO.DatZahTrg, &tradeType, &status, &tradeProposalDAO.RazlogOdbij,
			&tradeProposalDAO.IdMenadzerPos, &tradeProposalDAO.IdMenadzerPrim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusAndTypeForReading(status, tradeType, &tradeProposalDAO)
		tradeProposal := &model.TradeProposal{}
		tradeProposal.FromDAO(&tradeProposalDAO)

		tradeProposals = append(tradeProposals, *tradeProposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeProposals, nil
}

func (repo *tradeProposalRepository) GetByID(id int) (*model.TradeProposal, error) {
	var tradeProposalDAO model.TradeProposalDAO
	var tradeType, status string
	row := repo.db.QueryRow("SELECT * FROM ZahtevZaTrgovinu WHERE IDZAHTRG = :1", id)
	if err := row.Scan(&tradeProposalDAO.IdZahTrg, &tradeProposalDAO.DatZahTrg, &tradeType, &status, &tradeProposalDAO.RazlogOdbij,
		&tradeProposalDAO.IdMenadzerPos, &tradeProposalDAO.IdMenadzerPrim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromStatusAndTypeForReading(status, tradeType, &tradeProposalDAO)
	tradeProposal := &model.TradeProposal{}
	tradeProposal.FromDAO(&tradeProposalDAO)

	return tradeProposal, nil
}

func (repo *tradeProposalRepository) GetAllReceivedByManagerID(managerID int) ([]model.TradeProposal, error) {
	rows, err := repo.db.Query("SELECT * FROM ZahtevZaTrgovinu WHERE IDMENADZERPRIM = :1", managerID) // TODO: MORAM DODATI DA SE NE UZMU OTKAZANA
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeProposals []model.TradeProposal
	for rows.Next() {
		var tradeProposalDAO model.TradeProposalDAO
		var tradeType, status string
		if err := rows.Scan(&tradeProposalDAO.IdZahTrg, &tradeProposalDAO.DatZahTrg, &tradeType, &status, &tradeProposalDAO.RazlogOdbij,
			&tradeProposalDAO.IdMenadzerPos, &tradeProposalDAO.IdMenadzerPrim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusAndTypeForReading(status, tradeType, &tradeProposalDAO)
		tradeProposal := &model.TradeProposal{}
		tradeProposal.FromDAO(&tradeProposalDAO)

		tradeProposals = append(tradeProposals, *tradeProposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeProposals, nil
}

func (repo *tradeProposalRepository) GetAllSentByManagerID(managerID int) ([]model.TradeProposal, error) {
	rows, err := repo.db.Query("SELECT * FROM ZahtevZaTrgovinu WHERE IDMENADZERPOS = :1", managerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all trade proposals: %v", err)
	}
	defer rows.Close()

	var tradeProposals []model.TradeProposal
	for rows.Next() {
		var tradeProposalDAO model.TradeProposalDAO
		var tradeType, status string
		if err := rows.Scan(&tradeProposalDAO.IdZahTrg, &tradeProposalDAO.DatZahTrg, &tradeType, &status, &tradeProposalDAO.RazlogOdbij,
			&tradeProposalDAO.IdMenadzerPos, &tradeProposalDAO.IdMenadzerPrim); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusAndTypeForReading(status, tradeType, &tradeProposalDAO)
		tradeProposal := &model.TradeProposal{}
		tradeProposal.FromDAO(&tradeProposalDAO)

		tradeProposals = append(tradeProposals, *tradeProposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return tradeProposals, nil
}

func (repo *tradeProposalRepository) Create(tradeProposal *model.TradeProposal) error {
	status, tradeType := fromStatusAndTypeForWriting(tradeProposal)
	_, err := repo.db.Exec("INSERT INTO ZahtevZaTrgovinu (IDZAHTRG, DATZAHTRG, TIPZAHTRG, STATUSZAHTRG, RAZLOGODBIJ, "+
		"IDMENADZERPOS, IDMENADZERPRIM) VALUES (:1, :2, :3, :4, :5, :6, :7)", tradeProposal.ID, tradeProposal.Date,
		tradeType, status, tradeProposal.DenialReason, tradeProposal.SenderId, tradeProposal.ReceiverId)
	if err != nil {
		return fmt.Errorf("failed to create a trade proposal: %v", err)
	}
	return nil
}

func (repo *tradeProposalRepository) Update(tradeProposal *model.TradeProposal) error {
	status, _ := fromStatusAndTypeForWriting(tradeProposal)
	_, err := repo.db.Exec("UPDATE ZahtevZaTrgovinu SET STATUSZAHTRG = :1, RAZLOGODBIJ = :2 WHERE IDZAHTRG = :3",
		status, tradeProposal.DenialReason, tradeProposal.ID)
	if err != nil {
		return fmt.Errorf("failed to update trade proposal: %v", err)
	}
	return nil
}

func (repo *tradeProposalRepository) GetLatest() (*model.TradeProposal, error) {
	var tradeProposalDAO model.TradeProposalDAO
	var tradeType, status string
	row := repo.db.QueryRow(`SELECT * 
  								   FROM ZahtevZaTrgovinu 
								   WHERE ROWNUM = 1 
								   ORDER BY IDZAHTRG DESC`) // The latest one will have the highest id value because of the sequencer created on the server side
	if err := row.Scan(&tradeProposalDAO.IdZahTrg, &tradeProposalDAO.DatZahTrg, &tradeType, &status, &tradeProposalDAO.RazlogOdbij,
		&tradeProposalDAO.IdMenadzerPos, &tradeProposalDAO.IdMenadzerPrim); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromStatusAndTypeForReading(status, tradeType, &tradeProposalDAO)
	tradeProposal := &model.TradeProposal{}
	tradeProposal.FromDAO(&tradeProposalDAO)

	return tradeProposal, nil
}

func fromStatusAndTypeForWriting(tradeProposal *model.TradeProposal) (string, string) {
	var status, tradeType string

	switch tradeProposal.Status {
	case 0:
		status = "IN_PROGRESS"
	case 1:
		status = "ACCEPTED"
	case 2:
		status = "DECLINED"
	default:
		status = "CANCELLED"
	}

	switch tradeProposal.Type {
	case 0:
		tradeType = "PLAYER_PLAYER"
	case 1:
		tradeType = "PLAYER_PICK"
	default:
		tradeType = "PICK_PICK"
	}

	return status, tradeType
}

func fromStatusAndTypeForReading(status string, tradeType string, tradeProposal *model.TradeProposalDAO) {
	switch status {
	case "IN_PROGRESS":
		tradeProposal.StatusZahTrg = 0
	case "ACCEPTED":
		tradeProposal.StatusZahTrg = 1
	case "DECLINED":
		tradeProposal.StatusZahTrg = 2
	default:
		tradeProposal.StatusZahTrg = 3
	}

	switch tradeType {
	case "PLAYER_PLAYER":
		tradeProposal.TipZahTrg = 0
	case "PLAYER_PICK":
		tradeProposal.TipZahTrg = 1
	default:
		tradeProposal.TipZahTrg = 2
	}

	return
}
