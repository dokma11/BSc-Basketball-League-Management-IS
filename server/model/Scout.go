package model

import "time"

type Scout struct {
	Employee
}

func newScout(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Scout, error) {
	scout := &Scout{
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
	}

	if err := scout.Validate(); err != nil {
		return nil, err
	}

	return scout, nil
}

func (s *Scout) Validate() error {
	err := s.Employee.Validate()
	if err != nil {
		return err
	}

	return nil
}
