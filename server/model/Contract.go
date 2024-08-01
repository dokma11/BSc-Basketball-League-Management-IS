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
	ID             int64
	SigningDate    time.Time
	ExpirationDate time.Time
	Value          string // (in millions)
	Option         ContractOption
	TeamId         int64 // Team foreign key
	ContractTypeId int64 // Contract type foreign key
}

func NewContract(idUgo int64, datPotUgo time.Time, datVazUgo time.Time, vredUgo string, opcUgo ContractOption) (*Contract, error) {
	contract := &Contract{
		ID:             idUgo,
		SigningDate:    datPotUgo,
		ExpirationDate: datVazUgo,
		Value:          vredUgo,
		Option:         opcUgo,
	}

	if err := contract.Validate(); err != nil {
		return nil, err
	}

	return contract, nil
}

func (c *Contract) Validate() error {
	if c.Value == "" {
		return errors.New("value field can't be empty")
	}
	if c.Option < 0 || c.Option > 2 {
		return errors.New("option field is invalid")
	}
	return nil
}

type ContractDAO struct {
	IdUgo     int64
	DatPotUgo time.Time // Signing date
	DatVazUgo time.Time // Expiration date
	VredUgo   string    // Value (in millions)
	OpcUgo    ContractOption
	IdTim     int64 // Team foreign key
	IdTipUgo  int64 // Contract type foreign key
}

func (c *Contract) FromDAO(contractDAO *ContractDAO) {
	c.ID = contractDAO.IdUgo
	c.SigningDate = contractDAO.DatPotUgo
	c.ExpirationDate = contractDAO.DatVazUgo
	c.Value = contractDAO.VredUgo
	c.Option = contractDAO.OpcUgo
	c.TeamId = contractDAO.IdTim
	c.ContractTypeId = contractDAO.IdTipUgo
}

type ContractResponseDTO struct {
	IdUgo     int64          `json:"idUgo"`
	DatPotUgo time.Time      `json:"datPotUgo"` // Signing date
	DatVazUgo time.Time      `json:"datVazUgo"` // Expiration date
	VredUgo   string         `json:"vredUgo"`   // Value (in millions)
	OpcUgo    ContractOption `json:"opcUgo"`
	IdTim     int64          `json:"idTim"`    // Team foreign key
	IdTipUgo  int64          `json:"idTipUgo"` // Contract type foreign key
}

func (c *Contract) FromModel(contractDTO *ContractResponseDTO) {
	contractDTO.IdUgo = c.ID
	contractDTO.DatPotUgo = c.SigningDate
	contractDTO.DatVazUgo = c.ExpirationDate
	contractDTO.VredUgo = c.Value
	contractDTO.OpcUgo = c.Option
	contractDTO.IdTim = c.TeamId
	contractDTO.IdTipUgo = c.ContractTypeId
}
