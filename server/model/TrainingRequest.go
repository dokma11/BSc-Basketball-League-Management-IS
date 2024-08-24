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
	ID                 int64
	OccurrenceDateTime time.Time
	OccurrenceLocation string
	Status             TrainingRequestStatus
	DenialReason       string
	CoachId            int64 // Coach foreign key
	Duration           string
	TrainingTypeName   string
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
	TrajTrng      string                `json:"trajTrng"`      // Duration
	NazTipTrng    string                `json:"nazTipTrng"`    // Training type name
}

func (t *TrainingRequest) FromDAO(trainingRequestDAO *TrainingRequestDAO) {
	t.ID = trainingRequestDAO.IdPozTrng
	t.OccurrenceDateTime = trainingRequestDAO.DatVrePozTrng
	t.OccurrenceLocation = trainingRequestDAO.MesOdrPozTrng
	t.Status = trainingRequestDAO.StatusPozTrng
	t.DenialReason = trainingRequestDAO.RazOdbPozTrng
	t.CoachId = trainingRequestDAO.IdTrener
	t.Duration = trainingRequestDAO.TrajTrng
	t.TrainingTypeName = trainingRequestDAO.NazTipTrng
}

type TrainingRequestResponseDTO struct {
	IdPozTrng     int64                 `json:"idPozTrng"`
	DatVrePozTrng time.Time             `json:"datVrePozTrng"` // Occurrence date and time
	MesOdrPozTrng string                `json:"mesOdrPozTrng"` // Occurrence location
	StatusPozTrng TrainingRequestStatus `json:"statusPozTrng"`
	RazOdbPozTrng string                `json:"razOdbPozTrng"` // Denial reason
	IdTrener      int64                 `json:"idTrener"`      // Coach foreign key
	TrajTrng      string                `json:"trajTrng"`      // Duration
	NazTipTrng    string                `json:"nazTipTrng"`    // Training type name
	IdRegrut      int64                 `json:"idRegrut"`      // Recruit foreign key
}

func (t *TrainingRequest) FromModel(trainingRequestDTO *TrainingRequestResponseDTO) {
	trainingRequestDTO.IdPozTrng = t.ID
	trainingRequestDTO.DatVrePozTrng = t.OccurrenceDateTime
	trainingRequestDTO.MesOdrPozTrng = t.OccurrenceLocation
	trainingRequestDTO.StatusPozTrng = t.Status
	trainingRequestDTO.RazOdbPozTrng = t.DenialReason
	trainingRequestDTO.IdTrener = t.CoachId
	trainingRequestDTO.TrajTrng = t.Duration
	trainingRequestDTO.NazTipTrng = t.TrainingTypeName
}

type TrainingRequestCreateDTO struct {
	DatVrePozTrng time.Time `json:"datVrePozTrng"` // Occurrence date and time
	MesOdrPozTrng string    `json:"mesOdrPozTrng"` // Occurrence location
	IdTrener      int64     `json:"idTrener"`      // Coach foreign key
	TrajTrng      string    `json:"trajTrng"`      // Duration
	NazTipTrng    string    `json:"nazTipTrng"`    // Training type name
	IdRegrut      int64     `json:"idRegrut"`      // Recruit foreign key
}

func (t *TrainingRequest) FromDTO(trainingRequestDTO *TrainingRequestCreateDTO) {
	t.OccurrenceDateTime = trainingRequestDTO.DatVrePozTrng
	t.OccurrenceLocation = trainingRequestDTO.MesOdrPozTrng
	t.Status = 0
	t.CoachId = trainingRequestDTO.IdTrener
	t.Duration = trainingRequestDTO.TrajTrng
	t.TrainingTypeName = trainingRequestDTO.NazTipTrng
}

type TrainingRequestUpdateDTO struct {
	IdPozTrng     int64                 `json:"idPozTrng"`
	StatusPozTrng TrainingRequestStatus `json:"statusPozTrng"`
	RazOdbPozTrng string                `json:"razOdbPozTrng"` // Denial reason
	IdRegrut      int64                 `json:"idRegrut"`      // Recruit foreign key
}

func (t *TrainingRequest) FromUpdateDTO(trainingRequestDTO *TrainingRequestUpdateDTO) {
	t.ID = trainingRequestDTO.IdPozTrng
	t.Status = trainingRequestDTO.StatusPozTrng
	t.DenialReason = trainingRequestDTO.RazOdbPozTrng
}
