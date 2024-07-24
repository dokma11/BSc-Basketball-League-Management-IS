package model

import (
	"errors"
	"time"
)

type TrainingRequestStatus int

const (
	PENDING TrainingRequestStatus = iota
	APPROVED
	DISAPPROVED
)

type TrainingRequest struct {
	IdPozTrng     int64                 `json:"idPozTrng"`
	DatVrePozTrng time.Time             `json:"datVrePozTrng"` // Occurrence date and time
	MesOdrPozTrng string                `json:"mesOdrPozTrng"` // Occurrence location
	StatusPozTrng TrainingRequestStatus `json:"statusPozTrng"`
	RazOdbPozTrng string                `json:"razOdbPozTrng"` // Denial reason
	IdTrener      int64                 `json:"idTrener"`      // Coach foreign key
}

func NewTrainingRequest(idPozTrng int64, datVrePozTrng time.Time, mesOdrPozTrng string, statusPozTrng TrainingRequestStatus,
	razOdbPozTrng string) (*TrainingRequest, error) {
	trainingRequest := &TrainingRequest{
		IdPozTrng:     idPozTrng,
		DatVrePozTrng: datVrePozTrng,
		MesOdrPozTrng: mesOdrPozTrng,
		StatusPozTrng: statusPozTrng,
		RazOdbPozTrng: razOdbPozTrng,
	}

	if err := trainingRequest.Validate(); err != nil {
		return nil, err
	}

	return trainingRequest, nil
}

func (t *TrainingRequest) Validate() error {
	if t.MesOdrPozTrng == "" {
		return errors.New("location field is empty")
	}
	if t.StatusPozTrng < 0 || t.StatusPozTrng > 2 {
		return errors.New("status field is invalid")
	}
	if t.RazOdbPozTrng == "" {
		return errors.New("denial reason field is empty")
	}
	return nil
}
