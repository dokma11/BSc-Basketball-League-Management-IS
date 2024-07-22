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
	IdPozInt     int64                  `json:"idInt"`
	MesOdrPozInt string                 `json:"mesOdrPozInt"` // Occurrence location
	DatVrePozInt time.Time              `json:"datVrePozInt"` // Occurrence date and time
	StatusPozInt InterviewRequestStatus `json:"statusPozInt"`
	RazOdbPozInt string                 `json:"razOdbPozInt"`
}

func NewInterviewRequest(idInt int64, mesOdrPozInt string, datVrePozInt time.Time, statusPozInt InterviewRequestStatus,
	razOdbPozInt string) (*InterviewRequest, error) {
	interviewRequest := &InterviewRequest{
		IdPozInt:     idInt,
		MesOdrPozInt: mesOdrPozInt,
		DatVrePozInt: datVrePozInt,
		StatusPozInt: statusPozInt,
		RazOdbPozInt: razOdbPozInt,
	}

	if err := interviewRequest.Validate(); err != nil {
		return nil, err
	}

	return interviewRequest, nil
}

func (i *InterviewRequest) Validate() error {
	if i.MesOdrPozInt == "" {
		return errors.New("location field is empty")
	}
	if i.StatusPozInt < 0 || i.StatusPozInt > 2 {
		return errors.New("status field is invalid")
	}
	if i.RazOdbPozInt == "" {
		return errors.New("denial reason field is empty")
	}
	return nil
}
