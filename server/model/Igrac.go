package model

import (
	"errors"
	"time"
)

type Igrac struct {
	Zaposleni
	VisIgr string   `json:"VisIgr"`
	TezIgr string   `json:"TezIgr"`
	PozIgr Pozicija `json:"PozIgr"`
}

func NewIgrac(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, visIgr string, tezIgr string, pozIgr Pozicija) (*Igrac, error) {
	igrac := &Igrac{
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
		VisIgr: visIgr,
		TezIgr: tezIgr,
		PozIgr: pozIgr,
	}

	return igrac, nil
}

func (i *Igrac) Validate() error {
	err := i.Zaposleni.Validate()
	if err != nil {
		return err
	}
	if i.VisIgr == "" {
		return errors.New("height field is empty")
	}
	if i.TezIgr == "" {
		return errors.New("weight field is empty")
	}
	if i.PozIgr < 0 || i.PozIgr > 4 {
		return errors.New("position field is invalid")
	}

	return nil
}
