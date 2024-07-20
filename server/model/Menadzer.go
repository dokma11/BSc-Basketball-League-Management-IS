package model

import "time"

type Menadzer struct {
	Zaposleni
}

func newMenadzer(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Menadzer, error) {
	menadzer := &Menadzer{
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

	return menadzer, nil
}

func (m *Menadzer) Validate() error {
	err := m.Zaposleni.Validate()
	if err != nil {
		return err
	}

	return nil
}
