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

// Oversees Gerund
type Oversees struct {
	IdSkaut int64 `json:"idSkaut"`
	IdReg   int64 `json:"idReg"`
	IdTrng  int64 `json:"idTrng"`
}
