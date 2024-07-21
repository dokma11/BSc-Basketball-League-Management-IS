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

type Korisnik struct {
	Id      int64     `json:"Id"`
	Email   string    `json:"Email"`
	Ime     string    `json:"Ime"`
	Prezime string    `json:"Prezime"`
	DatRodj time.Time `json:"DatRodj"`
	Lozinka string    `json:"Lozinka"`
	Uloga   Uloga     `json:"Uloga"`
}

func NewKorisnik(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga) (*Korisnik, error) {
	korisnik := &Korisnik{
		Id:      id,
		Email:   email,
		Ime:     ime,
		Prezime: prezime,
		DatRodj: datRodj,
		Lozinka: lozinka,
		Uloga:   uloga,
	}

	return korisnik, nil
}

func (k *Korisnik) Validate() error {
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
