package model

import (
	"errors"
	"time"
)

type Recruit struct {
	User
	PhoneNumber  string
	Height       string
	Weight       string
	Position     Pozicija
	AverageRank  string
	AverageGrade string
	DraftId      int64
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
	KonTelefonReg string
	VisReg        string
	TezReg        string
	PozReg        Pozicija
	ProsRankReg   string
	ProsOcReg     string
	IdDraft       int64
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

type RecruitResponseDTO struct {
	UserResponseDTO
	KonTelefonReg string   `json:"konTelefonReg"` // Phone number
	VisReg        string   `json:"visReg"`        // Height
	TezReg        string   `json:"tezReg"`        // Weight
	PozReg        Pozicija `json:"pozReg"`
	ProsRankReg   string   `json:"prosRankReg"` // Average rank (ESPN)
	ProsOcReg     string   `json:"prosOcReg"`   // Average grade (ESPN)
	IdDraft       int64    `json:"idDraft"`     // Draft foreign key
}

func (r *Recruit) FromModel(recruitDTO *RecruitResponseDTO) {
	recruitDTO.UserResponseDTO.Id = r.User.ID
	recruitDTO.UserResponseDTO.Email = r.User.Email
	recruitDTO.UserResponseDTO.Ime = r.User.FirstName
	recruitDTO.UserResponseDTO.Prezime = r.User.LastName
	recruitDTO.UserResponseDTO.DatRodj = r.User.DateOfBirth
	recruitDTO.UserResponseDTO.Lozinka = r.User.Password
	recruitDTO.UserResponseDTO.Uloga = r.User.Role
	recruitDTO.KonTelefonReg = r.PhoneNumber
	recruitDTO.VisReg = r.Height
	recruitDTO.TezReg = r.Weight
	recruitDTO.PozReg = r.Position
	recruitDTO.ProsRankReg = r.AverageRank
	recruitDTO.ProsOcReg = r.AverageGrade
	recruitDTO.IdDraft = r.DraftId
}

type RecruitCreateDTO struct {
	ID            int64    `json:"id"`
	KonTelefonReg string   `json:"konTelefonReg"` // Phone number
	VisReg        string   `json:"visReg"`        // Height
	TezReg        string   `json:"tezReg"`        // Weight
	PozReg        Pozicija `json:"pozReg"`        // Position
	ProsRankReg   string   `json:"prosRankReg"`   // Average rank (ESPN)
	ProsOcReg     string   `json:"prosOcReg"`     // Average grade (ESPN)
}

func (r *Recruit) FromDTO(recruitDTO *RecruitCreateDTO) {
	r.ID = recruitDTO.ID
	r.PhoneNumber = recruitDTO.KonTelefonReg
	r.Height = recruitDTO.VisReg
	r.Weight = recruitDTO.TezReg
	r.Position = recruitDTO.PozReg
	r.AverageRank = recruitDTO.ProsRankReg
	r.AverageGrade = recruitDTO.ProsOcReg
}
