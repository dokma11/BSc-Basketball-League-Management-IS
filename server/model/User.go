package model

import (
	"errors"
	"time"
)

type Uloga int

const (
	UloRegrut    Uloga = iota // Recruit role
	UloZaposleni              // Employee role
)

type Pozicija int

const (
	PG Pozicija = iota // Point Guard
	SG                 // Shooting Guard
	SF                 // Small Forward
	PF                 // Power Forward
	C                  // Center
)

type User struct {
	ID          int64
	Email       string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Password    string
	Role        Uloga
}

func NewUser(id int64, email string, ime string, prezime string, datRodj time.Time,
	lozinka string, uloga Uloga) (*User, error) {
	user := &User{
		ID:          id,
		Email:       email,
		FirstName:   ime,
		LastName:    prezime,
		DateOfBirth: datRodj,
		Password:    lozinka,
		Role:        uloga,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.New("name field is empty")
	}
	if u.LastName == "" {
		return errors.New("surname field is empty")
	}
	if u.Email == "" {
		return errors.New("email field is empty")
	}
	if u.Password == "" {
		return errors.New("password field is empty")
	}
	if u.Role < 0 || u.Role > 1 {
		return errors.New("role field is invalid")
	}
	return nil
}

type UserDAO struct {
	Id      int64
	Email   string
	Ime     string    // First name
	Prezime string    // Last name
	DatRodj time.Time // Date of birth
	Lozinka string    // Password
	Uloga   Uloga     // Role
}

func (u *User) FromDAO(userDAO *UserDAO) {
	u.ID = userDAO.Id
	u.Email = userDAO.Email
	u.FirstName = userDAO.Ime
	u.LastName = userDAO.Prezime
	u.DateOfBirth = userDAO.DatRodj
	u.Password = userDAO.Lozinka
	u.Role = userDAO.Uloga
}

type UserResponseDTO struct {
	Id      int64     `json:"id"`
	Email   string    `json:"email"`
	Ime     string    `json:"ime"`     // First name
	Prezime string    `json:"prezime"` // Last name
	DatRodj time.Time `json:"datRodj"` // Date of birth
	Lozinka string    `json:"lozinka"` // Password
	Uloga   Uloga     `json:"uloga"`   // Role
}

func (u *User) FromModel(userDTO *UserResponseDTO) {
	userDTO.Id = u.ID
	userDTO.Email = u.Email
	userDTO.Ime = u.FirstName
	userDTO.Prezime = u.LastName
	userDTO.DatRodj = u.DateOfBirth
	userDTO.Lozinka = u.Password
	userDTO.Uloga = u.Role
}

type UserUpdateDTO struct {
	Id      int64     `json:"id"`
	Email   string    `json:"email"`
	Ime     string    `json:"ime"`     // First name
	Prezime string    `json:"prezime"` // Last name
	DatRodj time.Time `json:"datRodj"` // Date of birth
	Lozinka string    `json:"lozinka"` // Password
	Uloga   Uloga     `json:"uloga"`   // Role
}

func (u *User) FromDTO(userDTO *UserUpdateDTO) {
	u.ID = userDTO.Id
	u.Email = userDTO.Email
	u.FirstName = userDTO.Ime
	u.LastName = userDTO.Prezime
	u.DateOfBirth = userDTO.DatRodj
	u.Password = userDTO.Lozinka
	u.Role = userDTO.Uloga
}
