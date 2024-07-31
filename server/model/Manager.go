package model

import "time"

type Manager struct {
	Employee
}

func newManager(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Manager, error) {
	manager := &Manager{
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

	if err := manager.Validate(); err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Validate() error {
	err := m.Employee.Validate()
	if err != nil {
		return err
	}

	return nil
}
