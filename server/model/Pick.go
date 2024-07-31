package model

import (
	"errors"
	"fmt"
	"strconv"
)

type Pick struct {
	ID        int64  `json:"idPik"`
	Order     string `json:"redBrPik"`   // Pick order
	Round     string `json:"brRunPik"`   // Pick round (can be first and second)
	Year      string `json:"godPik"`     // Pick year
	ManagerId int64  `json:"idMenadzer"` // Manager that used the pick foreign key
	TeamId    int64  `json:"idTim"`      // Team foreign key
}

func NewPick(idPik int64, redBrPik string, brRunPik string, godPik string) (*Pick, error) {
	pick := &Pick{
		ID:    idPik,
		Order: redBrPik,
		Round: brRunPik,
		Year:  godPik,
	}

	if err := pick.Validate(); err != nil {
		return nil, err
	}

	return pick, nil
}

func (p *Pick) Validate() error {
	pickOrder, err := strconv.Atoi(p.Order)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if pickOrder > 30 || pickOrder < 1 { // There are 30 picks in each round
		return errors.New("pick order number field is invalid")
	}
	pickRound, err := strconv.Atoi(p.Round)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if pickRound > 2 || pickRound < 1 { // There are only two rounds
		return errors.New("pick round number field is invalid")
	}
	if len(p.Year) != 4 {
		return errors.New("pick year field is invalid")
	}
	return nil
}

// Bira Gerund
type Bira struct {
	IdRegrut int64 `json:"idRegrut"` // Recruit foreign key
	IdPik    int64 `json:"idPik"`    // Pick foreign key
}

type PickDAO struct {
	IdPik      int64  `json:"idPik"`
	RedBrPik   string `json:"redBrPik"`   // Pick order
	BrRunPik   string `json:"brRunPik"`   // Pick round (can be first and second)
	GodPik     string `json:"godPik"`     // Pick year
	IdMenadzer int64  `json:"idMenadzer"` // Manager that used the pick foreign key
	IdTim      int64  `json:"idTim"`      // Team foreign key
}

func (p *Pick) FromDAO(pickDAO *PickDAO) {
	p.ID = pickDAO.IdPik
	p.Order = pickDAO.RedBrPik
	p.Round = pickDAO.BrRunPik
	p.Year = pickDAO.GodPik
	p.ManagerId = pickDAO.IdMenadzer
	p.TeamId = pickDAO.IdTim
}
