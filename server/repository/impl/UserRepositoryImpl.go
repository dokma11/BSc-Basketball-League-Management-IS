package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) GetAll() ([]model.User, error) {
	rows, err := repo.db.Query("SELECT * FROM User")
	if err != nil {
		return nil, fmt.Errorf("failed to query all users: %v", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
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

func (repo *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	row := repo.db.QueryRow("SELECT * FROM User WHERE ID = :1", id)
	if err := row.Scan(&user.Id, &user.Ime, &user.Prezime, &user.Email, &user.DatRodj, &user.Lozinka, &user.Uloga); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &user, nil
}

func (repo *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	row := repo.db.QueryRow("SELECT * FROM User WHERE EMAIL = :1", email)
	if err := row.Scan(&user.Id, &user.Ime, &user.Prezime, &user.Email, &user.DatRodj, &user.Lozinka, &user.Uloga); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &user, nil
}
