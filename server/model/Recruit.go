package model

import (
	"errors"
	"time"
)

type Recruit struct {
	User
	KonTelefonReg string   `json:"KonTelefonReg"`
	MesRodjReg    string   `json:"MesRodjReg"`
	VisReg        string   `json:"VisReg"`
	TezReg        string   `json:"TezReg"`
	PozReg        Pozicija `json:"PozReg"`
	ProsRankReg   string   `json:"ProsRankReg"`
	ProsOcReg     string   `json:"ProsOcReg"`
}

func NewRecruit(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga, konTelefonReg string, mesRodjReg string, visReg string,
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
		MesRodjReg:    mesRodjReg,
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
	if r.MesRodjReg == "" {
		return errors.New("birth place field is empty")
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
