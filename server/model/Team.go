package model

import (
	"errors"
	"time"
)

type Team struct {
	ID                int64
	Name              string
	EstablishmentYear string
	Location          string
}

func NewTeam(idTeam int64, nazTeam string, godOsnTeam string, lokTeam string) (*Team, error) {
	team := &Team{
		ID:                idTeam,
		Name:              nazTeam,
		EstablishmentYear: godOsnTeam,
		Location:          lokTeam,
	}

	if err := team.Validate(); err != nil {
		return nil, err
	}

	return team, nil
}

func (t *Team) Validate() error {
	if t.Name == "" {
		return errors.New("name field was not set")
	}
	if t.Location == "" {
		return errors.New("location field was not set")
	}
	if len(t.EstablishmentYear) != 4 {
		return errors.New("foundation date field must have at least 4 digits")
	}
	return nil
}

type TeamDAO struct {
	IdTim     int64
	NazTim    string // Name
	GodOsnTim string // Establishment year
	LokTim    string // Location
}

func (t *Team) FromDAO(teamDAO *TeamDAO) {
	t.ID = teamDAO.IdTim
	t.Name = teamDAO.NazTim
	t.EstablishmentYear = teamDAO.GodOsnTim
	t.Location = teamDAO.LokTim
}

type TeamResponseDTO struct {
	IdTim     int64  `json:"idTim"`
	NazTim    string `json:"nazTim"`    // Name
	GodOsnTim string `json:"godOsnTim"` // Establishment year
	LokTim    string `json:"lokTim"`    // Location
}

func (t *Team) FromModel(teamDTO *TeamResponseDTO) {
	teamDTO.IdTim = t.ID
	teamDTO.NazTim = t.Name
	teamDTO.GodOsnTim = t.EstablishmentYear
	teamDTO.LokTim = t.Location
}

type WishlistAsset struct {
	IdZeljTim     int64     `json:"idZeljTim"`
	DatDodZeljTim time.Time `json:"datDodZeljTim"` // Date of creation
	BelesZeljTim  string    `json:"belesZeljTim"`  // Notes
	IdTipZelje    int64     `json:"idTipZelje"`    // Wishlist Asset Type foreign key
	IdPrava       int64     `json:"idPrava"`       // Draft Rights foreign key
	IdPik         int64     `json:"idPik"`         // Pick foreign key
	IdIgrac       int64     `json:"idIgrac"`       // Player foreign key
	IdTim         int64     `json:"idTim"`         // Team foreign key
	IdRegrut      int64     `json:"idRegrut"`      // Recruit foreign key
}

type WishlistAssetType struct {
	IdTipZelje  int64  `json:"idTipZelje"`
	NazTipZelje string `json:"nazTipZelje"`
}
