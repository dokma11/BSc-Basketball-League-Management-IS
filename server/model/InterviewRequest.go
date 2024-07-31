package model

import (
	"errors"
	"time"
)

type InterviewRequestStatus int

const (
	WAITING InterviewRequestStatus = iota
	AFFIRMED
	REJECTED
)

type InterviewRequest struct {
	ID                 int64                  `json:"idInt"`
	OccurrenceLocation string                 `json:"mesOdrPozInt"` // Occurrence location
	OccurrenceDateTime time.Time              `json:"datVrePozInt"` // Occurrence date and time
	Status             InterviewRequestStatus `json:"statusPozInt"`
	DenialReason       string                 `json:"razOdbPozInt"` // Denial reason
	RecruitId          int64                  `json:"idRegrut"`     // Recruit foreign key
	CoachId            int64                  `json:"idTrenter"`    // Coach foreign key
}

func NewInterviewRequest(idInt int64, mesOdrPozInt string, datVrePozInt time.Time, statusPozInt InterviewRequestStatus,
	razOdbPozInt string) (*InterviewRequest, error) {
	interviewRequest := &InterviewRequest{
		ID:                 idInt,
		OccurrenceLocation: mesOdrPozInt,
		OccurrenceDateTime: datVrePozInt,
		Status:             statusPozInt,
		DenialReason:       razOdbPozInt,
	}

	if err := interviewRequest.Validate(); err != nil {
		return nil, err
	}

	return interviewRequest, nil
}

func (i *InterviewRequest) Validate() error {
	if i.OccurrenceLocation == "" {
		return errors.New("location field is empty")
	}
	if i.Status < 0 || i.Status > 2 {
		return errors.New("status field is invalid")
	}
	if i.DenialReason == "" {
		return errors.New("denial reason field is empty")
	}
	return nil
}

type InterviewRequestDAO struct {
	IdPozInt     int64                  `json:"idInt"`
	MesOdrPozInt string                 `json:"mesOdrPozInt"` // Occurrence location
	DatVrePozInt time.Time              `json:"datVrePozInt"` // Occurrence date and time
	StatusPozInt InterviewRequestStatus `json:"statusPozInt"`
	RazOdbPozInt string                 `json:"razOdbPozInt"` // Denial reason
	IdRegrut     int64                  `json:"idRegrut"`     // Recruit foreign key
	IdTrener     int64                  `json:"idTrenter"`    // Coach foreign key
}

func (i *InterviewRequest) FromDAO(interviewRequestDAO *InterviewRequestDAO) {
	i.ID = interviewRequestDAO.IdPozInt
	i.OccurrenceLocation = interviewRequestDAO.MesOdrPozInt
	i.OccurrenceDateTime = interviewRequestDAO.DatVrePozInt
	i.Status = interviewRequestDAO.StatusPozInt
	i.DenialReason = interviewRequestDAO.RazOdbPozInt
	i.RecruitId = interviewRequestDAO.IdRegrut
	i.CoachId = interviewRequestDAO.IdTrener
}
