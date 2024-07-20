package model

import "time"

type Skaut struct {
	Zaposleni
}

func newSkaut(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Skaut, error) {
	skaut := &Skaut{
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
	}

	return skaut, nil
}

func (s *Skaut) Validate() error {
	err := s.Zaposleni.Validate()
	if err != nil {
		return err
	}

	return nil
}
