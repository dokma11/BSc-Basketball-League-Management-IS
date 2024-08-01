package model

import "errors"

type TradeSubjectType int

const (
	Igrac         TradeSubjectType = iota // Player
	Pik                                   // Pick
	PravaNaIgraca                         // Draft Rights
)

type TradeSubject struct {
	ID              int64
	Type            TradeSubjectType
	DraftRightsId   int64 // Draft Rights foreign key
	PlayerId        int64 // Player foreign key
	TradeProposalId int64 // Trade Request foreign key
	PickId          int64 // Pick foreign key
}

func NewTradeSubject(idPredTrg int64, tipPredTrg TradeSubjectType) (*TradeSubject, error) {
	tradeSubject := &TradeSubject{
		ID:   idPredTrg,
		Type: tipPredTrg,
	}

	if err := tradeSubject.Validate(); err != nil {
		return nil, err
	}

	return tradeSubject, nil
}

func (t *TradeSubject) Validate() error {
	if t.Type < 0 || t.Type > 2 {
		return errors.New("trade subject type field is invalid")
	}
	return nil
}

type TradeSubjectDAO struct {
	IdPredTrg  int64
	TipPredTrg TradeSubjectType
	IdPrava    int64 // Draft Rights foreign key
	IdIgrac    int64 // Player foreign key
	IdZahTrg   int64 // Trade Request foreign key
	IdPik      int64 // Pick foreign key
}

func (t *TradeSubject) FromDAO(tradeSubjectDAO *TradeSubjectDAO) {
	t.ID = tradeSubjectDAO.IdPredTrg
	t.Type = tradeSubjectDAO.TipPredTrg
	t.DraftRightsId = tradeSubjectDAO.IdPrava
	t.PlayerId = tradeSubjectDAO.IdIgrac
	t.TradeProposalId = tradeSubjectDAO.IdZahTrg
	t.PickId = tradeSubjectDAO.IdPik
}

type TradeSubjectResponseDTO struct {
	IdPredTrg  int64            `json:"idPredTrg"`
	TipPredTrg TradeSubjectType `json:"tipPredTrg"`
	IdPrava    int64            `json:"idPrava"`  // Draft Rights foreign key
	IdIgrac    int64            `json:"idIgrac"`  // Player foreign key
	IdZahTrg   int64            `json:"idZahTrg"` // Trade Request foreign key
	IdPik      int64            `json:"idPik"`    // Pick foreign key
}

func (t *TradeSubject) FromModel(tradeSubjectDTO *TradeSubjectResponseDTO) {
	tradeSubjectDTO.IdPredTrg = t.ID
	tradeSubjectDTO.TipPredTrg = t.Type
	tradeSubjectDTO.IdPrava = t.DraftRightsId
	tradeSubjectDTO.IdIgrac = t.PlayerId
	tradeSubjectDTO.IdZahTrg = t.TradeProposalId
	tradeSubjectDTO.IdPik = t.PickId
}

type TradeSubjectCreateDTO struct {
	TipPredTrg TradeSubjectType `json:"tipPredTrg"`
	IdPrava    int64            `json:"idPrava"` // Draft Rights foreign key
	IdIgrac    int64            `json:"idIgrac"` // Player foreign key
	IdPik      int64            `json:"idPik"`   // Pick foreign key
}

type TradeSubjectDetailsResponseDTO struct {
	IdPredTrg  int64            `json:"idPredTrg"`
	TipPredTrg TradeSubjectType `json:"tipPredTrg"`
	IdPrava    int64            `json:"idPrava"`  // Draft Rights foreign key
	IdZahTrg   int64            `json:"idZahTrg"` // Trade Proposal foreign key
	IdIgrac    int64            `json:"idIgrac"`  // Player foreign key
	IdPik      int64            `json:"idPik"`    // Pick foreign key
	IdTim      int64            `json:"idTim"`    // Team foreign key
}
