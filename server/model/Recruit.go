package model

import (
	"errors"
	"time"
)

type Recruit struct {
	User
	KonTelefonReg string   `json:"konTelefonReg"`
	VisReg        string   `json:"visReg"`
	TezReg        string   `json:"tezReg"`
	PozReg        Pozicija `json:"pozReg"`
	ProsRankReg   string   `json:"prosRankReg"`
	ProsOcReg     string   `json:"prosOcReg"`
	IdDraft       int64    `json:"idDraft"`
}

func NewRecruit(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, konTelefonReg string, visReg string,
	tezReg string, pozReg Pozicija, prosRankReg string, prosOcReg string) (*Recruit, error) {
	recruit := &Recruit{
		User: User{
			Id:      id,
			Email:   email,
			Ime:     ime,
			Prezime: prezime,
			DatRodj: datRodj,
			Lozinka: lozinka,
			Uloga:   uloga,
		},
		KonTelefonReg: konTelefonReg,
		VisReg:        visReg,
		TezReg:        tezReg,
		PozReg:        pozReg,
		ProsRankReg:   prosRankReg,
		ProsOcReg:     prosOcReg,
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
	if r.KonTelefonReg == "" {
		return errors.New("phone number field is empty")
	}
	if r.VisReg == "" {
		return errors.New("height field is empty")
	}
	if r.TezReg == "" {
		return errors.New("weight field is empty")
	}
	if r.PozReg < 0 || r.PozReg > 4 {
		return errors.New("position field is invalid")
	}
	if r.ProsRankReg == "" {
		return errors.New("average rank field is empty")
	}
	if r.ProsOcReg == "" {
		return errors.New("average grade field is empty")
	}
	return nil
}
