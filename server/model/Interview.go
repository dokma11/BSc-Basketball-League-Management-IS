package model

import (
	"errors"
	"time"
)

type Interview struct {
	ID                 int64     `json:"id"`
	OccurrenceLocation string    `json:"mesOdrInt"` // Occurrence location
	OccurrenceDateTime time.Time `json:"datVreInt"` // Occurrence date and time
	Notes              string    `json:"belesInt"`  // Notes taken from interview
	InterviewRequestId int64     `json:"idPozInt"`  // Interview request foreign key
	RecruitId          int64     `json:"idRegrut"`  // Recruit foreign key
}

func NewInterview(idInt int64, mesOdrInt string, datVreInt time.Time, belesInt string) (*Interview, error) {
	interview := &Interview{
		ID:                 idInt,
		OccurrenceLocation: mesOdrInt,
		OccurrenceDateTime: datVreInt,
		Notes:              belesInt,
	}

	if err := interview.Validate(); err != nil {
		return nil, err
	}

	return interview, nil
}

func (i *Interview) Validate() error {
	if i.OccurrenceLocation == "" {
		return errors.New("location field is empty")
	}
	if i.Notes == "" {
		return errors.New("notes field is empty")
	}
	return nil
}

type InterviewDAO struct {
	IdInt     int64     `json:"idInt"`
	MesOdrInt string    `json:"mesOdrInt"` // Occurrence location
	DatVreInt time.Time `json:"datVreInt"` // Occurrence date and time
	BelesInt  string    `json:"belesInt"`  // Notes taken from interview
	IdPozInt  int64     `json:"idPozInt"`  // Interview request foreign key
	IdRegrut  int64     `json:"idRegrut"`  // Recruit foreign key
}

func (i *Interview) FromDAO(interviewDAO *InterviewDAO) {
	i.ID = interviewDAO.IdInt
	i.OccurrenceLocation = interviewDAO.MesOdrInt
	i.OccurrenceDateTime = interviewDAO.DatVreInt
	i.Notes = interviewDAO.BelesInt
	i.InterviewRequestId = interviewDAO.IdPozInt
	i.RecruitId = interviewDAO.IdRegrut
}
