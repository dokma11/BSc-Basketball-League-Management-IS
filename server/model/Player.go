package model

import (
	"errors"
	"time"
)

type Player struct {
	Employee
	VisIgr string   `json:"VisIgr"` // Height
	TezIgr string   `json:"TezIgr"` // Weight
	PozIgr Pozicija `json:"PozIgr"` // Position (Point Guard, Shooting Guard, Small Forward, Power Forward, Center)
}

func NewPlayer(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, visIgr string, tezIgr string, pozIgr Pozicija) (*Player, error) {
	player := &Player{
		Employee: Employee{
			User: User{
				Id:      id,
				Email:   email,
				Ime:     ime,
				Prezime: prezime,
				DatRodj: datRodj,
				Lozinka: lozinka,
				Uloga:   uloga,
			},
			UloZap: ulogaZaposlenog,
			MbrZap: mbrZap,
		},
		VisIgr: visIgr,
		TezIgr: tezIgr,
		PozIgr: pozIgr,
	}

	if err := player.Validate(); err != nil {
		return nil, err
	}

	return player, nil
}

func (p *Player) Validate() error {
	err := p.Employee.Validate()
	if err != nil {
		return err
	}
	if p.VisIgr == "" {
		return errors.New("height field is empty")
	}
	if p.TezIgr == "" {
		return errors.New("weight field is empty")
	}
	if p.PozIgr < 0 || p.PozIgr > 4 {
		return errors.New("position field is invalid")
	}

	return nil
}
