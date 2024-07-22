package model

import "errors"

type Draft struct {
	IdDraft     int64  `json:"idDraft"`
	GodOdrDraft string `json:"godOdrDraft"`
	LokOdrDraft string `json:"lokOdrDraft"`
}

func NewDraft(idDraft int64, godOdrDraft string, lokOdrDraft string) (*Draft, error) {
	draft := &Draft{
		IdDraft:     idDraft,
		GodOdrDraft: godOdrDraft,
		LokOdrDraft: lokOdrDraft,
	}

	if err := draft.Validate(); err != nil {
		return nil, err
	}

	return draft, nil
}

func (d *Draft) Validate() error {
	if d.GodOdrDraft == "" {
		return errors.New("occurrence year field was not set")
	}
	if d.LokOdrDraft == "" {
		return errors.New("location field was not set")
	}

	return nil
}
