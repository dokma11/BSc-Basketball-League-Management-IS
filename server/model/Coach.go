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
	YearsOfExperience string              `json:"yearsOfExperience"`
	Specialization    CoachSpecialization `json:"specialization"`
}

func NewCoach(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, godIskTrener string, specTrener CoachSpecialization) (*Coach, error) {
	coach := &Coach{
		Employee: Employee{
			User: User{
				ID:          id,
				Email:       email,
				FirstName:   ime,
				LastName:    prezime,
				DateOfBirth: datRodj,
				Password:    lozinka,
				Role:        uloga,
			},
			Role:                 ulogaZaposlenog,
			IdentificationNumber: mbrZap,
		},
		YearsOfExperience: godIskTrener,
		Specialization:    specTrener,
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
	if c.YearsOfExperience == "" {
		return errors.New("years of experience field is empty")
	}
	if c.Specialization < 0 || c.Specialization > 2 {
		return errors.New("specialization field is invalid")
	}
	return nil
}

type CoachDAO struct {
	Employee
	GodIskTrener string              // Years of experience
	SpecTrener   CoachSpecialization // Coach specialization
}

func (c *Coach) FromDAO(coachDAO *CoachDAO) {
	c.Employee = coachDAO.Employee
	c.YearsOfExperience = coachDAO.GodIskTrener
	c.Specialization = coachDAO.SpecTrener
}
