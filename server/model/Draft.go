package model

import "errors"

type Draft struct {
	ID                 int64
	OccurrenceYear     string
	OccurrenceLocation string
}

func NewDraft(idDraft int64, godOdrDraft string, lokOdrDraft string) (*Draft, error) {
	draft := &Draft{
		ID:                 idDraft,
		OccurrenceYear:     godOdrDraft,
		OccurrenceLocation: lokOdrDraft,
	}

	if err := draft.Validate(); err != nil {
		return nil, err
	}

	return draft, nil
}

func (d *Draft) Validate() error {
	if d.OccurrenceYear == "" {
		return errors.New("occurrence year field was not set")
	}
	if d.OccurrenceLocation == "" {
		return errors.New("location field was not set")
	}
	return nil
}

type DraftDAO struct {
	IdDraft     int64
	GodOdrDraft string
	LokOdrDraft string
}

func (d *Draft) FromDAO(draftDAO *DraftDAO) {
	d.ID = draftDAO.IdDraft
	d.OccurrenceYear = draftDAO.GodOdrDraft
	d.OccurrenceLocation = draftDAO.LokOdrDraft
}

type DraftResponseDTO struct {
	IdDraft     int64  `json:"idDraft"`
	GodOdrDraft string `json:"godOdrDraft"`
	LokOdrDraft string `json:"lokOdrDraft"`
}

func (d *Draft) FromModel(draftDTO *DraftResponseDTO) {
	draftDTO.IdDraft = d.ID
	draftDTO.GodOdrDraft = d.OccurrenceYear
	draftDTO.LokOdrDraft = d.OccurrenceLocation
}
