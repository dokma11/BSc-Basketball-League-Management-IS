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
		var trade model.Trade
		var tradeType string
		if err := rows.Scan(&trade.IdTrg, &trade.DatTrg, &tradeType, &trade.IdZahTrg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		mapTradeEnumsForReading(tradeType, &trade)

		trades = append(trades, trade)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trades, nil
}

func (repo *tradeRepository) GetByID(id int) (*model.Trade, error) {
	var trade model.Trade
	var tradeType string
	row := repo.db.QueryRow("SELECT * FROM TRGOVINA WHERE IDTRG = :1", id)
	if err := row.Scan(&trade.IdTrg, &trade.DatTrg, &tradeType, &trade.IdZahTrg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	mapTradeEnumsForReading(tradeType, &trade)

	return &trade, nil
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
		var trade model.Trade
		var tradeType string
		if err := rows.Scan(&trade.IdTrg, &trade.DatTrg, &tradeType, &trade.IdZahTrg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		mapTradeEnumsForReading(tradeType, &trade)

		trades = append(trades, trade)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trades, nil
}

func (repo *tradeRepository) Create(trade *model.Trade) error {
	tradeType := mapTradeEnumsForWriting(trade)
	_, err := repo.db.Exec("INSERT INTO TRGOVINA (IDTRG, DATTRG, TIPTRG, IDZAHTRG) VALUES (:1, :2, :3, :4)",
		trade.IdTrg, trade.DatTrg, tradeType, trade.IdZahTrg)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create a trade: %v", err)
	}
	return nil
}

func mapTradeEnumsForWriting(trade *model.Trade) string {
	var tradeType string

	switch trade.TipTrg {
	case 0:
		tradeType = "PLAYER_PLAYER"
	case 1:
		tradeType = "PLAYER_PICK"
	default:
		tradeType = "PICK_PICK"
	}

	return tradeType
}

func mapTradeEnumsForReading(tradeType string, trade *model.Trade) {
	switch tradeType {
	case "PLAYER_PLAYER":
		trade.TipTrg = 0
	case "PLAYER_PICK":
		trade.TipTrg = 1
	default:
		trade.TipTrg = 2
	}
	return
}
