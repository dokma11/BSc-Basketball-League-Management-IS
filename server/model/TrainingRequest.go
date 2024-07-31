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
	ID                 int64                 `json:"idPozTrng"`
	OccurrenceDateTime time.Time             `json:"datVrePozTrng"` // Occurrence date and time
	OccurrenceLocation string                `json:"mesOdrPozTrng"` // Occurrence location
	Status             TrainingRequestStatus `json:"statusPozTrng"`
	DenialReason       string                `json:"razOdbPozTrng"` // Denial reason
	CoachId            int64                 `json:"idTrener"`      // Coach foreign key
}

func NewTrainingRequest(idPozTrng int64, datVrePozTrng time.Time, mesOdrPozTrng string, statusPozTrng TrainingRequestStatus,
	razOdbPozTrng string) (*TrainingRequest, error) {
	trainingRequest := &TrainingRequest{
		ID:                 idPozTrng,
		OccurrenceDateTime: datVrePozTrng,
		OccurrenceLocation: mesOdrPozTrng,
		Status:             statusPozTrng,
		DenialReason:       razOdbPozTrng,
	}

	if err := trainingRequest.Validate(); err != nil {
		return nil, err
	}

	return trainingRequest, nil
}

func (t *TrainingRequest) Validate() error {
	if t.OccurrenceLocation == "" {
		return errors.New("location field is empty")
	}
	if t.Status < 0 || t.Status > 2 {
		return errors.New("status field is invalid")
	}
	if t.DenialReason == "" {
		return errors.New("denial reason field is empty")
	}
	return nil
}

type TrainingRequestDAO struct {
	IdPozTrng     int64                 `json:"idPozTrng"`
	DatVrePozTrng time.Time             `json:"datVrePozTrng"` // Occurrence date and time
	MesOdrPozTrng string                `json:"mesOdrPozTrng"` // Occurrence location
	StatusPozTrng TrainingRequestStatus `json:"statusPozTrng"`
	RazOdbPozTrng string                `json:"razOdbPozTrng"` // Denial reason
	IdTrener      int64                 `json:"idTrener"`      // Coach foreign key
}

func (t *TrainingRequest) FromDAO(trainingRequestDAO *TrainingRequestDAO) {
	t.ID = trainingRequestDAO.IdPozTrng
	t.OccurrenceDateTime = trainingRequestDAO.DatVrePozTrng
	t.OccurrenceLocation = trainingRequestDAO.MesOdrPozTrng
	t.Status = trainingRequestDAO.StatusPozTrng
	t.DenialReason = trainingRequestDAO.RazOdbPozTrng
	t.CoachId = trainingRequestDAO.IdTrener
}
