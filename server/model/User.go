package model

import (
	"errors"
	"time"
)

type Uloga int

const ( // Moram ovde pripaziti
	UloRegrut    Uloga = iota // Recruit role
	UloZaposleni              // Employee role
)

type Pozicija int

const (
	PG Pozicija = iota // Point Guard
	SG                 // Shooting Guard
	SF                 // Small Forward
	PF                 // Power Forward
	C                  // Center
)

type User struct {
	Id      int64     `json:"id"`
	Email   string    `json:"email"`
	Ime     string    `json:"ime"`
	Prezime string    `json:"prezime"`
	DatRodj time.Time `json:"datRodj"`
	Lozinka string    `json:"lozinka"`
	Uloga   Uloga     `json:"uloga"`
}

func NewUser(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga) (*User, error) {
	user := &User{
		Id:      id,
		Email:   email,
		Ime:     ime,
		Prezime: prezime,
		DatRodj: datRodj,
		Lozinka: lozinka,
		Uloga:   uloga,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (k *User) Validate() error {
	if k.Ime == "" {
		return errors.New("name field is empty")
	}
	if k.Prezime == "" {
		return errors.New("surname field is empty")
	}
	if k.Email == "" {
		return errors.New("email field is empty")
	}
	if k.Lozinka == "" {
		return errors.New("password field is empty")
	}
	if k.Uloga < 0 || k.Uloga > 1 {
		return errors.New("role field is invalid")
	}

	return nil
}
