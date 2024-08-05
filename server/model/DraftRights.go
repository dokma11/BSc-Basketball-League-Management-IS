package model

type DraftRight struct {
	ID          int64
	TeamId      int64 // Team foreign key (team that is in the possession of this)
	RecruitId   int64 // Recruit foreign key
	PickId      int64 // Pick foreign key (pick that was used to pick the recruit)
	Untouchable bool  // Is the draft right added to the untouchables list
	Tradeable   bool  // Is the draft right added to the trade list
}

func NewDraftRight(idPrava int64, untouchable bool, tradeable bool) (*DraftRight, error) {
	draftRight := &DraftRight{
		ID:          idPrava,
		Untouchable: untouchable,
		Tradeable:   tradeable,
	}
	return draftRight, nil
}

type DraftRightDAO struct {
	IdPrava        int64
	IdTim          int64  // Team foreign key (team that is in the possession of this)
	IdRegrut       int64  // Recruit foreign key
	IdPik          int64  // Pick foreign key (pick that was used to pick the recruit)
	NedodListPrava string // Is the draft right added to the untouchables list
	TrgListPrava   string // Is the draft right added to the trade list
}

func (dr *DraftRight) FromDAO(draftRightDAO *DraftRightDAO) {
	dr.ID = draftRightDAO.IdPrava
	dr.TeamId = draftRightDAO.IdTim
	dr.RecruitId = draftRightDAO.IdRegrut
	dr.PickId = draftRightDAO.IdPik
	if draftRightDAO.NedodListPrava == "TRUE" {
		dr.Untouchable = true
	} else if draftRightDAO.NedodListPrava == "FALSE" {
		dr.Untouchable = false
	}
	if draftRightDAO.TrgListPrava == "TRUE" {
		dr.Tradeable = true
	} else if draftRightDAO.TrgListPrava == "FALSE" {
		dr.Tradeable = false
	}
}

type DraftRightResponseDTO struct {
	IdPrava        int64 `json:"idPrava"`
	IdTim          int64 `json:"idTim"`          // Team foreign key (team that is in the possession of this)
	IdRegrut       int64 `json:"idRegrut"`       // Recruit foreign key
	IdPik          int64 `json:"idPik"`          // Pick foreign key (pick that was used to pick the recruit)
	NedodListPrava bool  `json:"nedodListPrava"` // Is the draft right added to the untouchables list
	TrgListPrava   bool  `json:"trgListPrava"`   // Is the draft right added to the trade list
}

func (dr *DraftRight) FromModel(draftRightDTO *DraftRightResponseDTO) {
	draftRightDTO.IdPrava = dr.ID
	draftRightDTO.IdTim = dr.TeamId
	draftRightDTO.IdRegrut = dr.RecruitId
	draftRightDTO.IdPik = dr.PickId
	draftRightDTO.NedodListPrava = dr.Untouchable
	draftRightDTO.TrgListPrava = dr.Tradeable
}

type DraftRightUpdateDTO struct {
	IdPrava        int64 `json:"idPrava"`
	NedodListPrava bool  `json:"nedodListPrava"` // Is the draft right added to the untouchables list
	TrgListPrava   bool  `json:"trgListPrava"`   // Is the draft right added to the trade list
}

func (dr *DraftRight) FromUpdateDTO(draftRightDTO *DraftRightUpdateDTO) {
	dr.ID = draftRightDTO.IdPrava
	dr.Untouchable = draftRightDTO.NedodListPrava
	dr.Tradeable = draftRightDTO.TrgListPrava
}
