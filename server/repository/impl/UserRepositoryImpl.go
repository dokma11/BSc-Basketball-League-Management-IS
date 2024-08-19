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
	rows, err := repo.db.Query("SELECT * FROM KORISNIK")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to query all users: %v", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var userDAO model.UserDAO
		var role string
		if err := rows.Scan(&userDAO.Id, &userDAO.Email, &userDAO.Ime, &userDAO.Prezime, &userDAO.DatRodj, &userDAO.Lozinka, &role); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRole(role, &userDAO)
		user := &model.User{}
		user.FromDAO(&userDAO)

		users = append(users, *user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

func (repo *userRepository) GetByID(id int) (*model.User, error) {
	var userDAO model.UserDAO
	var role string
	row := repo.db.QueryRow("SELECT * FROM KORISNIK WHERE ID = :1", id)
	if err := row.Scan(&userDAO.Id, &userDAO.Email, &userDAO.Ime, &userDAO.Prezime, &userDAO.DatRodj, &userDAO.Lozinka, &role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRole(role, &userDAO)
	user := &model.User{}
	user.FromDAO(&userDAO)

	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*model.User, error) {
	var userDAO model.UserDAO
	var role string
	row := repo.db.QueryRow("SELECT * FROM KORISNIK WHERE EMAIL = :1", email)
	if err := row.Scan(&userDAO.Id, &userDAO.Email, &userDAO.Ime, &userDAO.Prezime, &userDAO.DatRodj, &userDAO.Lozinka, &role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRole(role, &userDAO)
	user := &model.User{}
	user.FromDAO(&userDAO)

	return user, nil
}

func (repo *userRepository) Update(user *model.User) error {
	_, err := repo.db.Exec("UPDATE KORISNIK SET IME = :1, PREZIME = :2, EMAIL = :3, DATRODJ = :4 WHERE IDUGO = :5",
		user.FirstName, user.LastName, user.Email, user.DateOfBirth, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update suer: %v", err)
	}
	return nil
}

func (repo *userRepository) Create(user *model.User) error {
	_, err := repo.db.Exec("INSERT INTO KORISNIK VALUES (0, :1, :2, :3, :4, :5, 'Regrut')",
		user.Email, user.FirstName, user.LastName, user.DateOfBirth, user.Password)
	if err != nil {
		return fmt.Errorf("failed to update suer: %v", err)
	}
	return nil
}

func fromRole(role string, userDAO *model.UserDAO) {
	if role == "Regrut" {
		userDAO.Uloga = 0
	} else if role == "Zaposleni" {
		userDAO.Uloga = 1
	}
}
