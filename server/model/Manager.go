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
