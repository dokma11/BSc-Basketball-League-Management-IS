package model

import (
	"errors"
	"time"
)

type Training struct {
	ID                 int64     `json:"idTrng"`
	Duration           string    `json:"trajTrng"`   // Duration
	OccurrenceDateTime time.Time `json:"datVreTrng"` // Occurrence date and time
	OccurrenceLocation string    `json:"mesOdrTrng"` // Occurrence location
	Notes              string    `json:"belesTrng"`  // Notes from training
	TrainingTypeId     int64     `json:"idTipTrng"`  // Training type foreign key
	TrainingRequestId  int64     `json:"idPozTrng"`  // Training Request foreign key
}

func NewTraining(idTrng int64, trajTrng string, datVreTrng time.Time, mesOdrTrng string, belesTrng string) (*Training, error) {
	training := &Training{
		ID:                 idTrng,
		Duration:           trajTrng,
		OccurrenceDateTime: datVreTrng,
		OccurrenceLocation: mesOdrTrng,
		Notes:              belesTrng,
	}

	if err := training.Validate(); err != nil {
		return nil, err
	}

	return training, nil
}

func (t *Training) Validate() error {
	if t.Duration == "" {
		return errors.New("duration field is empty")
	}
	if t.OccurrenceLocation == "" {
		return errors.New("location field is empty")
	}
	if t.Notes == "" {
		return errors.New("notes field is empty")
	}
	return nil
}

type TrainingDAO struct {
	IdTrng     int64     `json:"idTrng"`
	TrajTrng   string    `json:"trajTrng"`   // Duration
	DatVreTrng time.Time `json:"datVreTrng"` // Occurrence date and time
	MesOdrTrng string    `json:"mesOdrTrng"` // Occurrence location
	BelesTrng  string    `json:"belesTrng"`  // Notes from training
	IdTipTrng  int64     `json:"idTipTrng"`  // Training type foreign key
	IdPozTrng  int64     `json:"idPozTrng"`  // Training Request foreign key
}

func (t *Training) FromDAO(trainingDAO *TrainingDAO) {
	t.ID = trainingDAO.IdTrng
	t.Duration = trainingDAO.TrajTrng
	t.OccurrenceDateTime = trainingDAO.DatVreTrng
	t.OccurrenceLocation = trainingDAO.MesOdrTrng
	t.Notes = trainingDAO.BelesTrng
	t.TrainingTypeId = trainingDAO.IdTipTrng
	t.TrainingRequestId = trainingDAO.IdPozTrng
}
