package model

import (
	"errors"
	"time"
)

type Team struct {
	ID                int64  `json:"idTim"`
	Name              string `json:"nazTim"`    // Name
	EstablishmentYear string `json:"godOsnTim"` // Establishment year
	Location          string `json:"lokTim"`    // Location
	// TODO: Razmotriti za liste samo da li FK da bude ili da pravim materijalizovani pogled
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
	IdTim     int64  `json:"idTim"`
	NazTim    string `json:"nazTim"`    // Name
	GodOsnTim string `json:"godOsnTim"` // Establishment year
	LokTim    string `json:"lokTim"`    // Location
	// TODO: Razmotriti za liste samo da li FK da bude ili da pravim materijalizovani pogled
}

func (t *Team) FromDAO(teamDAO *TeamDAO) {
	t.ID = teamDAO.IdTim
	t.Name = teamDAO.NazTim
	t.EstablishmentYear = teamDAO.GodOsnTim
	t.Location = teamDAO.LokTim
}

type AssetForTrade struct {
	IdImoTrgTim     int64     `json:"idImoTrgTim"`
	DatDodImoTrgTim time.Time `json:"datDodImoTrgTim"`
	BelesImoTrgTim  string    `json:"belesImoTrgTim"`
	IdTipImoTrg     int64     `json:"idTipImoTrg"`
	IdTim           *int64    `json:"idTim"`
}

type AssetForTradeType struct {
	IdTipImoTrg  int64  `json:"idTipImoTrg"`
	NazTipImoTrg string `json:"nazTipImoTrg"`
}

type UntouchableAsset struct {
	IdNedIdTim      int64     `json:"idNedIdTim"`
	DatDodNedImoTim time.Time `json:"datDodNedImoTim"`
	BelesNedImoTim  string    `json:"belesNedImoTim"`
	IdTim           *int64    `json:"idTim"`
}

type UntouchableAssetType struct {
	IdTipNedImo  int64  `json:"idTipNedImo"`
	NazTipNedImo string `json:"nazTipNedImo"`
}

type WishlistAsset struct {
	IdZeljTim     int64     `json:"idZeljTim"`
	DatDodZeljTim time.Time `json:"datDodZeljTim"`
	BelesZeljTim  *string   `json:"belesZeljTim"`
	IdTipZelje    int64     `json:"idTipZelje"`
	IdTim         *int64    `json:"idTim"`
}

type WishlistAssetType struct {
	IdTipZelje  int64  `json:"idTipZelje"`
	NazTipZelje string `json:"nazTipZelje"`
}
