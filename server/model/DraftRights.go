package model

import "errors"

type DraftRight struct {
	IdPrava          int64    `json:"idPrava"`
	ImeIgrPrava      string   `json:"imeIgrPrava"`      // Name of the player included in the draft rights
	PrezimeIgrPrava  string   `json:"prezimeIgrPrava"`  // Surname of the player included in the draft rights
	PozicijaIgrPrava Pozicija `json:"pozicijaIgrPrava"` // Position of the player included in the draft rights
	StatusPrava      string   `json:"statusPrava"`      // Status (can be used or unused)
}

func NewDraftRight(idPrava int64, imeIgrPrava string, prezimeIgrPrava string, pozicijaIgrPrava Pozicija,
	statusPrava string) (*DraftRight, error) {
	draftRight := &DraftRight{
		IdPrava:          idPrava,
		ImeIgrPrava:      imeIgrPrava,
		PrezimeIgrPrava:  prezimeIgrPrava,
		PozicijaIgrPrava: pozicijaIgrPrava,
		StatusPrava:      statusPrava,
	}

	if err := draftRight.Validate(); err != nil {
		return nil, err
	}

	return draftRight, nil
}

func (dr *DraftRight) Validate() error {
	if dr.ImeIgrPrava == "" {
		return errors.New("name field is empty")
	}
	if dr.PrezimeIgrPrava == "" {
		return errors.New("surname field is empty")
	}
	if dr.PozicijaIgrPrava < 0 || dr.PozicijaIgrPrava > 4 {
		return errors.New("position field is invalid")
	}
	if dr.StatusPrava == "" {
		return errors.New("status field is empty")
	}

	return nil
}
