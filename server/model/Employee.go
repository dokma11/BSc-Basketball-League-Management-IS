package model

import (
	"errors"
	"time"
)

type UlogaZaposlenog int

const (
	UlogaMenadzer UlogaZaposlenog = iota // Moram ovde pripaziti
	UlogaIgrac
	UlogaTrener
	UlogaSkaut
)

type Employee struct {
	User
	UloZap UlogaZaposlenog `json:"uloZap"`
	MbrZap string          `json:"mbrZap"`
}

func NewEmployee(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Employee, error) {
	employee := &Employee{
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
	}

	if err := employee.Validate(); err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *Employee) Validate() error {
	err := e.User.Validate()
	if err != nil {
		return err
	}
	if e.UloZap < 0 || e.UloZap > 3 {
		return errors.New("employee role field is invalid")
	}
	if e.MbrZap == "" {
		return errors.New("identification number field is empty")
	}

	return nil
}
