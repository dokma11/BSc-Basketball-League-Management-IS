package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type zaposleniRepository struct {
	db *sql.DB
}

func NewZaposleniRepository(db *sql.DB) repository.ZaposleniRepository {
	return &zaposleniRepository{db}
}

func (repo *zaposleniRepository) GetAll() ([]model.Zaposleni, error) {
	rows, err := repo.db.Query("SELECT * FROM ZAPOSLENI")
	if err != nil {
		return nil, fmt.Errorf("failed to query all employees: %v", err)
	}
	defer rows.Close()

	var employees []model.Zaposleni
	for rows.Next() {
		var employee model.Zaposleni
		if err := rows.Scan(&employee.Id, &employee.Ime, &employee.Prezime, &employee.Email, &employee.DatRodj,
			&employee.Lozinka, &employee.Uloga, &employee.UloZap, &employee.MbrZap); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return employees, nil
}

func (repo *zaposleniRepository) GetByID(id int) (*model.Zaposleni, error) {
	var employee model.Zaposleni
	row := repo.db.QueryRow("SELECT * FROM ZAPOSLENI WHERE ID = :1", id)
	if err := row.Scan(&employee.Id, &employee.Ime, &employee.Prezime, &employee.Email, &employee.DatRodj,
		&employee.Lozinka, &employee.Uloga, &employee.UloZap, &employee.MbrZap); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &employee, nil
}
