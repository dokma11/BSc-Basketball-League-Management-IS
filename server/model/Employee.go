package model

import (
	"errors"
	"time"
)

type UlogaZaposlenog int

const (
	Menadzer UlogaZaposlenog = iota // Moram ovde pripaziti
	UlogaIgrac
	Trener
	Skaut
)

type Employee struct {
	User
	Role                 UlogaZaposlenog `json:"role"`
	IdentificationNumber string          `json:"identificationNumber"`
	ContractId           int64           `json:"contractId"` // Contract foreign key
}

func NewEmployee(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, ulogaZaposlenog UlogaZaposlenog, mbrZap string) (*Employee, error) {
	employee := &Employee{
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
	}

	if err := employee.Validate(); err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *Employee) Validate() error {
	err := e.User.Validate()
	if err != nil {
		return err
	}
	if e.Role < 0 || e.Role > 3 {
		return errors.New("employee role field is invalid")
	}
	if e.IdentificationNumber == "" {
		return errors.New("identification number field is empty")
	}
	return nil
}

type EmployeeDAO struct {
	User
	UloZap UlogaZaposlenog `json:"uloZap"`
	MbrZap string          `json:"mbrZap"` // Identification number
	IdUgo  int64           `json:"idUgo"`  // Contract foreign key
}

func (e *Employee) FromDAO(employeeDAO *EmployeeDAO) {
	e.User = employeeDAO.User
	e.Role = employeeDAO.UloZap
	e.IdentificationNumber = employeeDAO.MbrZap
	e.ContractId = employeeDAO.IdUgo
}
