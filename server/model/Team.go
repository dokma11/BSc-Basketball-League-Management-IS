package model

import "errors"

// Team Obratiti paznju na one liste neophodne
type Team struct {
	IdTim     int64  `json:"idTim"`
	NazTim    string `json:"nazTim"`
	GodOsnTim string `json:"godOsnTim"`
	LokTim    string `json:"lokTim"`
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
