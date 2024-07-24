package model

import "errors"

type DraftRight struct {
	IdPrava          int64    `json:"idPrava"`
	ImeIgrPrava      string   `json:"imeIgrPrava"`      // Name of the player included in the draft rights
	PrezimeIgrPrava  string   `json:"prezimeIgrPrava"`  // Surname of the player included in the draft rights
	PozicijaIgrPrava Pozicija `json:"pozicijaIgrPrava"` // Position of the player included in the draft rights
	IdTim            int64    `json:"idTim"`            // Team foreign key (team that is in the possession of this)
	IdRegrut         int64    `json:"idRegrut"`         // Recruit foreign key
	IdPik            int64    `json:"idPik"`            // Pick foreign key (pick that was used to pick the recruit)
}

func NewDraftRight(idPrava int64, imeIgrPrava string, prezimeIgrPrava string, pozicijaIgrPrava Pozicija) (*DraftRight, error) {
	draftRight := &DraftRight{
		IdPrava:          idPrava,
		ImeIgrPrava:      imeIgrPrava,
		PrezimeIgrPrava:  prezimeIgrPrava,
		PozicijaIgrPrava: pozicijaIgrPrava,
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
	return nil
}
