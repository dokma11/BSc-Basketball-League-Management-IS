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

		if tradeType == "PLAYER_PLAYER" {
			trade.TipTrg = 0
		} else if tradeType == "PLAYER_PICK" {
			trade.TipTrg = 1
		} else if tradeType == "PICK_PICK" {
			trade.TipTrg = 2
		}

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

	if tradeType == "PLAYER_PLAYER" {
		trade.TipTrg = 0
	} else if tradeType == "PLAYER_PICK" {
		trade.TipTrg = 1
	} else if tradeType == "PICK_PICK" {
		trade.TipTrg = 2
	}

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

		if tradeType == "PLAYER_PLAYER" {
			trade.TipTrg = 0
		} else if tradeType == "PLAYER_PICK" {
			trade.TipTrg = 1
		} else if tradeType == "PICK_PICK" {
			trade.TipTrg = 2
		}

		trades = append(trades, trade)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trades, nil
}

func (repo *tradeRepository) Create(trade *model.Trade) error {
	_, err := repo.db.Exec("INSERT INTO TRGOVINA (IDTRG, DATTRG, TIPTRG, IDZAHTRG) VALUES (:1, :2, :3, :4)",
		trade.IdTrg, trade.DatTrg, trade.TipTrg, trade.IdZahTrg)
	if err != nil {
		return fmt.Errorf("failed to create a trade: %v", err)
	}
	return nil
}
