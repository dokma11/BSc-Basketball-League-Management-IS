package model

import "errors"

// Tim Obratiti paznju na one liste neophodne
type Tim struct {
	IdTim     int64  `json:"IdTim"`
	NazTim    string `json:"NazTim"`
	GodOsnTim string `json:"GodOsnTim"`
	LokTim    string `json:"LokTim"`
}

func NewTim(idTim int64, nazTim string, godOsnTim string, lokTim string) (*Tim, error) {
	tim := &Tim{
		IdTim:     idTim,
		NazTim:    nazTim,
		GodOsnTim: godOsnTim,
		LokTim:    lokTim,
	}

	if err := tim.Validate(); err != nil {
		return nil, err
	}

	return tim, nil
}

func (tim *Tim) Validate() error {
	if tim.NazTim == "" {
		return errors.New("team name was not set")
	}
	if tim.LokTim == "" {
		return errors.New("team location was not set")
	}
	if len(tim.GodOsnTim) != 4 {
		return errors.New("foundation date must have at least 4 digits")
	}
	return nil
}
