package model

import (
	"errors"
	"fmt"
	"strconv"
)

type Pick struct {
	IdPik    int64  `json:"idPik"`
	RedBrPik string `json:"idPik"` // Pick order
	BrRunPik string `json:"idPik"` // Pick round (can be first and second)
	GodPik   string `json:"idPik"` // Pick year
}

func NewPick(idPik int64, redBrPik string, brRunPik string, godPik string) (*Pick, error) {
	pick := &Pick{
		IdPik:    idPik,
		RedBrPik: redBrPik,
		BrRunPik: brRunPik,
		GodPik:   godPik,
	}

	if err := pick.Validate(); err != nil {
		return nil, err
	}

	return pick, nil
}

func (p *Pick) Validate() error {
	pickOrder, err := strconv.Atoi(p.RedBrPik)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if pickOrder > 30 || pickOrder < 1 { // There are 30 picks in each round
		return errors.New("pick order number is invalid")
	}

	pickRound, err := strconv.Atoi(p.BrRunPik)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if pickRound > 2 || pickRound < 1 { // There are only two rounds
		return errors.New("pick round number is invalid")
	}

	if len(p.GodPik) != 4 {
		return errors.New("pick year is invalid")
	}

	return nil
}
