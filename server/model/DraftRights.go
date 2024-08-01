package model

import "errors"

type DraftRight struct {
	ID             int64
	PlayerName     string   // Name of the player included in the draft rights
	PlayerSurname  string   // Surname of the player included in the draft rights
	PlayerPosition Pozicija // Position of the player included in the draft rights
	TeamId         int64    // Team foreign key (team that is in the possession of this)
	RecruitId      int64    // Recruit foreign key
	PickId         int64    // Pick foreign key (pick that was used to pick the recruit)
}

func NewDraftRight(idPrava int64, imeIgrPrava string, prezimeIgrPrava string, pozicijaIgrPrava Pozicija) (*DraftRight, error) {
	draftRight := &DraftRight{
		ID:             idPrava,
		PlayerName:     imeIgrPrava,
		PlayerSurname:  prezimeIgrPrava,
		PlayerPosition: pozicijaIgrPrava,
	}

	if err := draftRight.Validate(); err != nil {
		return nil, err
	}

	return draftRight, nil
}

func (dr *DraftRight) Validate() error {
	if dr.PlayerName == "" {
		return errors.New("name field is empty")
	}
	if dr.PlayerSurname == "" {
		return errors.New("surname field is empty")
	}
	if dr.PlayerPosition < 0 || dr.PlayerPosition > 4 {
		return errors.New("position field is invalid")
	}
	return nil
}

type DraftRightDAO struct {
	IdPrava          int64
	ImeIgrPrava      string   // Name of the player included in the draft rights
	PrezimeIgrPrava  string   // Surname of the player included in the draft rights
	PozicijaIgrPrava Pozicija // Position of the player included in the draft rights
	IdTim            int64    // Team foreign key (team that is in the possession of this)
	IdRegrut         int64    // Recruit foreign key
	IdPik            int64    // Pick foreign key (pick that was used to pick the recruit)
}

func (dr *DraftRight) FromDAO(draftRightDAO *DraftRightDAO) {
	dr.ID = draftRightDAO.IdPrava
	dr.PlayerName = draftRightDAO.ImeIgrPrava
	dr.PlayerSurname = draftRightDAO.PrezimeIgrPrava
	dr.PlayerPosition = draftRightDAO.PozicijaIgrPrava
	dr.TeamId = draftRightDAO.IdTim
	dr.RecruitId = draftRightDAO.IdRegrut
	dr.PickId = draftRightDAO.IdPik
}

type DraftRightResponseDTO struct {
	IdPrava          int64    `json:"idPrava"`
	ImeIgrPrava      string   `json:"imeIgrPrava"`      // Name of the player included in the draft rights
	PrezimeIgrPrava  string   `json:"prezimeIgrPrava"`  // Surname of the player included in the draft rights
	PozicijaIgrPrava Pozicija `json:"pozicijaIgrPrava"` // Position of the player included in the draft rights
	IdTim            int64    `json:"idTim"`            // Team foreign key (team that is in the possession of this)
	IdRegrut         int64    `json:"idRegrut"`         // Recruit foreign key
	IdPik            int64    `json:"idPik"`            // Pick foreign key (pick that was used to pick the recruit)
}

func (dr *DraftRight) FromModel(draftRightDTO *DraftRightResponseDTO) {
	draftRightDTO.IdPrava = dr.ID
	draftRightDTO.ImeIgrPrava = dr.PlayerName
	draftRightDTO.PrezimeIgrPrava = dr.PlayerSurname
	draftRightDTO.PozicijaIgrPrava = dr.PlayerPosition
	draftRightDTO.IdTim = dr.TeamId
	draftRightDTO.IdRegrut = dr.RecruitId
	draftRightDTO.IdPik = dr.PickId
}
