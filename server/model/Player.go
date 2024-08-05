package model

import (
	"errors"
	"time"
)

type Player struct {
	Employee
	Height      string
	Weight      string
	Position    Pozicija
	Untouchable bool // Is the player added to the untouchables list
	Tradeable   bool // Is the player added to the trade list
}

func NewPlayer(id int64, email string, ime string, prezime string, datRodj time.Time, lozinka string, uloga Uloga,
	ulogaZaposlenog UlogaZaposlenog, mbrZap string, visIgr string, tezIgr string, pozIgr Pozicija) (*Player, error) {
	player := &Player{
		Employee: Employee{
			User: User{
				ID:          id,
				Email:       email,
				FirstName:   ime,
				LastName:    prezime,
				DateOfBirth: datRodj,
				Password:    lozinka,
				Role:        uloga,
			},
			Role:                 ulogaZaposlenog,
			IdentificationNumber: mbrZap,
		},
		Height:   visIgr,
		Weight:   tezIgr,
		Position: pozIgr,
	}

	if err := player.Validate(); err != nil {
		return nil, err
	}

	return player, nil
}

func (p *Player) Validate() error {
	err := p.Employee.Validate()
	if err != nil {
		return err
	}
	if p.Height == "" {
		return errors.New("height field is empty")
	}
	if p.Weight == "" {
		return errors.New("weight field is empty")
	}
	if p.Position < 0 || p.Position > 4 {
		return errors.New("position field is invalid")
	}
	return nil
}

type PlayerDAO struct {
	Employee
	VisIgr       string   // Height
	TezIgr       string   // Weight
	PozIgr       Pozicija // Position (Point Guard, Shooting Guard, Small Forward, Power Forward, Center)
	NedodListIgr string   // Is player added to the untouchables list
	TrgListIgr   string   // Is player added to the trade list
}

func (p *Player) FromDAO(playerDAO *PlayerDAO) {
	p.ID = playerDAO.ID
	p.Email = playerDAO.Email
	p.FirstName = playerDAO.FirstName
	p.LastName = playerDAO.LastName
	p.DateOfBirth = playerDAO.DateOfBirth
	p.Height = playerDAO.VisIgr
	p.Weight = playerDAO.TezIgr
	p.Position = playerDAO.PozIgr
	if playerDAO.NedodListIgr == "TRUE" {
		p.Untouchable = true
	} else if playerDAO.NedodListIgr == "FALSE" {
		p.Untouchable = false
	}
	if playerDAO.TrgListIgr == "TRUE" {
		p.Tradeable = true
	} else if playerDAO.TrgListIgr == "FALSE" {
		p.Tradeable = false
	}
}

type PlayerResponseDTO struct {
	EmployeeResponseDTO
	VisIgr       string   `json:"visIgr"`       // Height
	TezIgr       string   `json:"tezIgr"`       // Weight
	PozIgr       Pozicija `json:"pozIgr"`       // Position (Point Guard, Shooting Guard, Small Forward, Power Forward, Center)
	NedodListIgr bool     `json:"nedodListIgr"` // Is player added to the untouchables list
	TrgListIgr   bool     `json:"trgListIgr"`   // Is player added to the trade list
}

func (p *Player) FromModel(playerDTO *PlayerResponseDTO) {
	playerDTO.EmployeeResponseDTO.Id = p.ID
	playerDTO.EmployeeResponseDTO.UserResponseDTO.Ime = p.FirstName
	playerDTO.EmployeeResponseDTO.UserResponseDTO.Prezime = p.LastName
	playerDTO.EmployeeResponseDTO.UserResponseDTO.Email = p.Email
	playerDTO.EmployeeResponseDTO.UserResponseDTO.DatRodj = p.DateOfBirth
	playerDTO.VisIgr = p.Height
	playerDTO.TezIgr = p.Weight
	playerDTO.PozIgr = p.Position
	playerDTO.NedodListIgr = p.Untouchable
	playerDTO.TrgListIgr = p.Tradeable
}

type PlayerUpdateDTO struct {
	ID           int64 `json:"id"`
	NedodListIgr bool  `json:"nedodListIgr"` // Is player added to the untouchables list
	TrgListIgr   bool  `json:"trgListIgr"`   // Is player added to the trade list
}

func (p *Player) FromUpdateDTO(playerDTO *PlayerUpdateDTO) {
	p.ID = playerDTO.ID
	p.Untouchable = playerDTO.NedodListIgr
	p.Tradeable = playerDTO.TrgListIgr
}
