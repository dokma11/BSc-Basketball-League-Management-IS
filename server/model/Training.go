package model

import (
	"errors"
	"time"
)

type Training struct {
	IdTrng     int64     `json:"idTrng"`
	TrajTrng   string    `json:"trajTrng"`   // Duration
	DatVreTrng time.Time `json:"datVreTrng"` // Occurrence date and time
	MesOdrTrng string    `json:"mesOdrTrng"` // Occurrence location
	BelesTrng  string    `json:"belesTrng"`  // Notes from training
}

func NewTraining(idTrng int64, trajTrng string, datVreTrng time.Time, mesOdrTrng string, belesTrng string) (*Training, error) {
	training := &Training{
		IdTrng:     idTrng,
		TrajTrng:   trajTrng,
		DatVreTrng: datVreTrng,
		MesOdrTrng: mesOdrTrng,
		BelesTrng:  belesTrng,
	}

	if err := training.Validate(); err != nil {
		return nil, err
	}

	return training, nil
}

func (t *Training) Validate() error {
	if t.TrajTrng == "" {
		return errors.New("duration field is empty")
	}
	if t.MesOdrTrng == "" {
		return errors.New("location field is empty")
	}
	if t.BelesTrng == "" {
		return errors.New("notes field is empty")
	}
	return nil
}
