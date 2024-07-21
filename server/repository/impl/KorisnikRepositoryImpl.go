package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type korisnikRepository struct {
	db *sql.DB
}

func NewKorisnikRepository(db *sql.DB) repository.KorisnikRepository {
	return &korisnikRepository{db}
}

func (repo *korisnikRepository) GetAll() ([]model.Korisnik, error) {
	rows, err := repo.db.Query("SELECT * FROM KORISNIK")
	if err != nil {
		return nil, fmt.Errorf("failed to query all users: %v", err)
	}
	defer rows.Close()

	var users []model.Korisnik
	for rows.Next() {
		var user model.Korisnik
		if err := rows.Scan(&user.Id, &user.Ime, &user.Prezime, &user.Email, &user.DatRodj, &user.Lozinka, &user.Uloga); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

func (repo *korisnikRepository) GetByID(id int) (*model.Korisnik, error) {
	var user model.Korisnik
	row := repo.db.QueryRow("SELECT * FROM KORISNIK WHERE ID = :1", id)
	if err := row.Scan(&user.Id, &user.Ime, &user.Prezime, &user.Email, &user.DatRodj, &user.Lozinka, &user.Uloga); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &user, nil
}

func (repo *korisnikRepository) GetByEmail(email string) (*model.Korisnik, error) {
	var user model.Korisnik
	row := repo.db.QueryRow("SELECT * FROM KORISNIK WHERE EMAIL = :1", email)
	if err := row.Scan(&user.Id, &user.Ime, &user.Prezime, &user.Email, &user.DatRodj, &user.Lozinka, &user.Uloga); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &user, nil
}
