package model

import (
	"errors"
	"time"
)

type Trener struct {
	Zaposleni
	GodIskTrener string `json:"GodIskTrener"`
	SpecTrener   string `json:"SpecTrener"`
}

func NewTrener(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, godIskTrener string, specTrener string) (*Trener, error) {
	trener := &Trener{
		Zaposleni: Zaposleni{
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
		},
		GodIskTrener: godIskTrener,
		SpecTrener:   specTrener,
	}

	return trener, nil
}

func (t *Trener) Validate() error {
	err := t.Zaposleni.Validate()
	if err != nil {
		return err
	}
	if t.GodIskTrener == "" {
		return errors.New("years of experience field is empty")
	}
	if t.SpecTrener == "" {
		return errors.New("specialization field is empty")
	}

	return nil
}
