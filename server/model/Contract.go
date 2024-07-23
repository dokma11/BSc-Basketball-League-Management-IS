package model

import (
	"errors"
	"time"
)

type ContractOption int

const (
	PLAYER_OPTION ContractOption = iota
	TEAM_OPTION
	NO_OPTION
)

type Contract struct {
	IdUgo     int64          `json:"idUgo"`
	DatPotUgo time.Time      `json:"datPotUgo"` // Signing date
	DatVazUgo time.Time      `json:"datVazUgo"` // Expiration date
	VredUgo   string         `json:"vredUgo"`   // Value (in millions)
	OpcUgo    ContractOption `json:"opcUgo"`
}

func NewContract(idUgo int64, datPotUgo time.Time, datVazUgo time.Time, vredUgo string, opcUgo ContractOption) (*Contract, error) {
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
	if c.OpcUgo < 0 || c.OpcUgo > 2 {
		return errors.New("option field is invalid")
	}
	return nil
}
