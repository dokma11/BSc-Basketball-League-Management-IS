package model

import (
	"errors"
	"time"
)

type CoachSpecialization int

const (
	OFFENSE CoachSpecialization = iota
	DEFENSE
	PLAYER_MANAGEMENT
)

type Coach struct {
	Employee
	GodIskTrener string              `json:"godIskTrener"`
	SpecTrener   CoachSpecialization `json:"specTrener"`
}

func NewCoach(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, godIskTrener string, specTrener CoachSpecialization) (*Coach, error) {
	coach := &Coach{
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
		GodIskTrener: godIskTrener,
		SpecTrener:   specTrener,
	}

	if err := coach.Validate(); err != nil {
		return nil, err
	}

	return coach, nil
}

func (c *Coach) Validate() error {
	err := c.Employee.Validate()
	if err != nil {
		return err
	}
	if c.GodIskTrener == "" {
		return errors.New("years of experience field is empty")
	}
	if c.SpecTrener < 0 || c.SpecTrener > 2 {
		return errors.New("specialization field is invalid")
	}

	return nil
}
