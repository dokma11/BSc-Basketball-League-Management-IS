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

type Zaposleni struct {
	Korisnik
	UloZap UlogaZaposlenog `json:"UloZap"`
	MbrZap string          `json:"MbrZap"`
}

func NewZaposleni(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Zaposleni, error) {
	zaposleni := &Zaposleni{
		Korisnik: Korisnik{
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

	return zaposleni, nil
}

func (z *Zaposleni) Validate() error {
	err := z.Korisnik.Validate()
	if err != nil {
		return err
	}
	if z.UloZap < 0 || z.UloZap > 3 {
		return errors.New("employee role field is invalid")
	}
	if z.MbrZap == "" {
		return errors.New("identification number field is empty")
	}

	return nil
}
