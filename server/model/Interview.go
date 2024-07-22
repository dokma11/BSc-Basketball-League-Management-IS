package model

import (
	"errors"
	"time"
)

type Interview struct {
	IdInt     int64     `json:"idInt"`
	MesOdrInt string    `json:"mesOdrInt"` // Occurrence location
	DatVreInt time.Time `json:"datVreInt"` // Occurrence date and time
	BelesInt  string    `json:"belesInt"`  // Notes taken from interview
}

func NewInterview(idInt int64, mesOdrInt string, datVreInt time.Time, belesInt string) (*Interview, error) {
	interview := &Interview{
		IdInt:     idInt,
		MesOdrInt: mesOdrInt,
		DatVreInt: datVreInt,
		BelesInt:  belesInt,
	}

	if err := interview.Validate(); err != nil {
		return nil, err
	}

	return interview, nil
}

func (i *Interview) Validate() error {
	if i.MesOdrInt == "" {
		return errors.New("location field is empty")
	}
	if i.BelesInt == "" {
		return errors.New("notes field is empty")
	}
	return nil
}
