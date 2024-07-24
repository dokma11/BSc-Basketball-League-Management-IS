package model

import (
	"errors"
	"time"
)

// Team Obratiti paznju na one liste neophodne
type Team struct {
	IdTim     int64  `json:"idTim"`
	NazTim    string `json:"nazTim"`    // Name
	GodOsnTim string `json:"godOsnTim"` // Establishment year
	LokTim    string `json:"lokTim"`    // Location
	// TODO: Razmotriti za liste samo da li FK da bude ili da pravim materijalizovani pogled
}

func NewTeam(idTeam int64, nazTeam string, godOsnTeam string, lokTeam string) (*Team, error) {
	team := &Team{
		IdTim:     idTeam,
		NazTim:    nazTeam,
		GodOsnTim: godOsnTeam,
		LokTim:    lokTeam,
	}

	if err := team.Validate(); err != nil {
		return nil, err
	}

	return team, nil
}

func (team *Team) Validate() error {
	if team.NazTim == "" {
		return errors.New("name field was not set")
	}
	if team.LokTim == "" {
		return errors.New("location field was not set")
	}
	if len(team.GodOsnTim) != 4 {
		return errors.New("foundation date field must have at least 4 digits")
	}
	return nil
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
