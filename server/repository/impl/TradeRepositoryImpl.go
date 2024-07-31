package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type tradeRepository struct {
	db *sql.DB
}

func NewTradeRepository(db *sql.DB) repository.TradeRepository {
	return &tradeRepository{db}
}

func (repo *tradeRepository) GetAll() ([]model.Trade, error) {
	rows, err := repo.db.Query("SELECT * FROM TRGOVINA")
	if err != nil {
		return nil, fmt.Errorf("failed to query all trades: %v", err)
	}
	defer rows.Close()

	var trades []model.Trade
	for rows.Next() {
		var tradeDAO model.TradeDAO
		var tradeType string
		if err := rows.Scan(&tradeDAO.IdTrg, &tradeDAO.DatTrg, &tradeType, &tradeDAO.IdZahTrg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTypeForReading(tradeType, &tradeDAO)
		trade := &model.Trade{}
		trade.FromDAO(&tradeDAO)

		trades = append(trades, *trade)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trades, nil
}

func (repo *tradeRepository) GetByID(id int) (*model.Trade, error) {
	var tradeDAO model.TradeDAO
	var tradeType string
	row := repo.db.QueryRow("SELECT * FROM TRGOVINA WHERE IDTRG = :1", id)
	if err := row.Scan(&tradeDAO.IdTrg, &tradeDAO.DatTrg, &tradeType, &tradeDAO.IdZahTrg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromTypeForReading(tradeType, &tradeDAO)
	trade := &model.Trade{}
	trade.FromDAO(&tradeDAO)

	return trade, nil
}

func (repo *tradeRepository) GetAllByTeamID(teamID int) ([]model.Trade, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM TRGOVINA") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all trades: %v", err)
	}
	defer rows.Close()

	var trades []model.Trade
	for rows.Next() {
		var tradeDAO model.TradeDAO
		var tradeType string
		if err := rows.Scan(&tradeDAO.IdTrg, &tradeDAO.DatTrg, &tradeType, &tradeDAO.IdZahTrg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromTypeForReading(tradeType, &tradeDAO)
		trade := &model.Trade{}
		trade.FromDAO(&tradeDAO)

		trades = append(trades, *trade)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trades, nil
}

func (repo *tradeRepository) Create(trade *model.Trade) error {
	tradeType := fromTypeForWriting(trade)
	_, err := repo.db.Exec("INSERT INTO TRGOVINA (IDTRG, DATTRG, TIPTRG, IDZAHTRG) VALUES (:1, :2, :3, :4)",
		trade.ID, trade.OccurrenceDate, tradeType, trade.TradeProposalId)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create a trade: %v", err)
	}
	return nil
}

func fromTypeForWriting(trade *model.Trade) string {
	var tradeType string
	switch trade.Type {
	case 0:
		tradeType = "PLAYER_PLAYER"
	case 1:
		tradeType = "PLAYER_PICK"
	default:
		tradeType = "PICK_PICK"
	}
	return tradeType
}

func fromTypeForReading(tradeType string, tradeDAO *model.TradeDAO) {
	switch tradeType {
	case "PLAYER_PLAYER":
		tradeDAO.TipTrg = 0
	case "PLAYER_PICK":
		tradeDAO.TipTrg = 1
	default:
		tradeDAO.TipTrg = 2
	}
	return
}
