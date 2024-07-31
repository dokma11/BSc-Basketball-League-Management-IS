package model

import (
	"errors"
	"time"
)

type Recruit struct {
	User
	PhoneNumber  string   `json:"konTelefonReg"`
	Height       string   `json:"visReg"`
	Weight       string   `json:"tezReg"`
	Position     Pozicija `json:"pozReg"`
	AverageRank  string   `json:"prosRankReg"`
	AverageGrade string   `json:"prosOcReg"`
	DraftId      int64    `json:"idDraft"`
}

func NewRecruit(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, konTelefonReg string, visReg string,
	tezReg string, pozReg Pozicija, prosRankReg string, prosOcReg string) (*Recruit, error) {
	recruit := &Recruit{
		User: User{
			ID:          id,
			Email:       email,
			FirstName:   ime,
			LastName:    prezime,
			DateOfBirth: datRodj,
			Password:    lozinka,
			Role:        uloga,
		},
		PhoneNumber:  konTelefonReg,
		Height:       visReg,
		Weight:       tezReg,
		Position:     pozReg,
		AverageRank:  prosRankReg,
		AverageGrade: prosOcReg,
	}

	if err := recruit.Validate(); err != nil {
		return nil, err
	}

	return recruit, nil
}

func (r *Recruit) Validate() error {
	err := r.User.Validate()
	if err != nil {
		return err
	}
	if r.PhoneNumber == "" {
		return errors.New("phone number field is empty")
	}
	if r.Height == "" {
		return errors.New("height field is empty")
	}
	if r.Weight == "" {
		return errors.New("weight field is empty")
	}
	if r.Position < 0 || r.Position > 4 {
		return errors.New("position field is invalid")
	}
	if r.AverageRank == "" {
		return errors.New("average rank field is empty")
	}
	if r.AverageGrade == "" {
		return errors.New("average grade field is empty")
	}
	return nil
}

type RecruitDAO struct {
	User
	KonTelefonReg string   `json:"konTelefonReg"`
	VisReg        string   `json:"visReg"`
	TezReg        string   `json:"tezReg"`
	PozReg        Pozicija `json:"pozReg"`
	ProsRankReg   string   `json:"prosRankReg"`
	ProsOcReg     string   `json:"prosOcReg"`
	IdDraft       int64    `json:"idDraft"`
}

func (r *Recruit) FromDAO(recruitDAO *RecruitDAO) {
	r.User = recruitDAO.User
	r.PhoneNumber = recruitDAO.KonTelefonReg
	r.Height = recruitDAO.VisReg
	r.Weight = recruitDAO.TezReg
	r.Position = recruitDAO.PozReg
	r.AverageRank = recruitDAO.ProsRankReg
	r.AverageGrade = recruitDAO.ProsOcReg
	r.DraftId = recruitDAO.IdDraft
}
