package model

import (
	"errors"
	"time"
)

type Training struct {
	ID                 int64
	Duration           string
	OccurrenceDateTime time.Time
	OccurrenceLocation string
	Notes              string
	TrainingTypeId     int64 // Training type foreign key
	TrainingRequestId  int64 // Training Request foreign key
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
	IdTrng     int64
	TrajTrng   string    // Duration
	DatVreTrng time.Time // Occurrence date and time
	MesOdrTrng string    // Occurrence location
	BelesTrng  string    // Notes from training
	IdTipTrng  int64     // Training type foreign key
	IdPozTrng  int64     // Training Request foreign key
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

type TrainingResponseDTO struct {
	IdTrng     int64     `json:"idTrng"`
	TrajTrng   string    `json:"trajTrng"`   // Duration
	DatVreTrng time.Time `json:"datVreTrng"` // Occurrence date and time
	MesOdrTrng string    `json:"mesOdrTrng"` // Occurrence location
	BelesTrng  string    `json:"belesTrng"`  // Notes from training
	IdTipTrng  int64     `json:"idTipTrng"`  // Training type foreign key
	IdPozTrng  int64     `json:"idPozTrng"`  // Training Request foreign key
}

func (t *Training) FromModel(trainingDTO *TrainingResponseDTO) {
	trainingDTO.IdTrng = t.ID
	trainingDTO.TrajTrng = t.Duration
	trainingDTO.DatVreTrng = t.OccurrenceDateTime
	trainingDTO.MesOdrTrng = t.OccurrenceLocation
	trainingDTO.BelesTrng = t.Notes
	trainingDTO.IdTipTrng = t.TrainingTypeId
	trainingDTO.IdPozTrng = t.TrainingRequestId
}

type TrainingCreateDTO struct {
	TrajTrng   string    `json:"trajTrng"`   // Duration
	DatVreTrng time.Time `json:"datVreTrng"` // Occurrence date and time
	MesOdrTrng string    `json:"mesOdrTrng"` // Occurrence location
	BelesTrng  string    `json:"belesTrng"`  // Notes from training
	IdTipTrng  int64     `json:"idTipTrng"`  // Training type foreign key
	IdPozTrng  int64     `json:"idPozTrng"`  // Training Request foreign key
}

func (t *Training) FromDTO(trainingDTO *TrainingCreateDTO) {
	t.Duration = trainingDTO.TrajTrng
	t.OccurrenceDateTime = trainingDTO.DatVreTrng
	t.OccurrenceLocation = trainingDTO.MesOdrTrng
	t.Notes = trainingDTO.BelesTrng
	t.TrainingTypeId = trainingDTO.IdTipTrng
	t.TrainingRequestId = trainingDTO.IdPozTrng
}
