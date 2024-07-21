package model

import "errors"

// Team Obratiti paznju na one liste neophodne
type Team struct {
	IdTim     int64  `json:"IdTeam"`
	NazTim    string `json:"NazTeam"`
	GodOsnTim string `json:"GodOsnTeam"`
	LokTim    string `json:"LokTeam"`
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
		return errors.New("team name was not set")
	}
	if team.LokTim == "" {
		return errors.New("team location was not set")
	}
	if len(team.GodOsnTim) != 4 {
		return errors.New("foundation date must have at least 4 digits")
	}
	return nil
}
