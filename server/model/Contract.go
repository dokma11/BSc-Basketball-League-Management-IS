package model

import (
	"errors"
	"time"
)

type Contract struct {
	IdUgo     int64     `json:"idUgo"`
	DatPotUgo time.Time `json:"datPotUgo"` // Signing date
	DatVazUgo time.Time `json:"datVazUgo"` // Expiration date
	VredUgo   string    `json:"vredUgo"`   // Value (in millions)
	OpcUgo    string    `json:"opcUgo"`    // Option (can have player-option, team-option or none) OVDE TREBA PROVERITI
}

func NewContract(idUgo int64, datPotUgo time.Time, datVazUgo time.Time, vredUgo string, opcUgo string) (*Contract, error) {
	contract := &Contract{
		IdUgo:     idUgo,
		DatPotUgo: datPotUgo,
		DatVazUgo: datVazUgo,
		VredUgo:   vredUgo,
		OpcUgo:    opcUgo,
	}

	if err := contract.Validate(); err != nil {
		return nil, err
	}

	return contract, nil
}

func (c *Contract) Validate() error {
	if c.VredUgo == "" {
		return errors.New("value field can't be empty")
	}
	if c.OpcUgo == "" {
		return errors.New("option field can't be empty")
	}

	return nil
}
