package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) repository.EmployeeRepository {
	return &employeeRepository{db}
}

func (repo *employeeRepository) GetAll() ([]model.Employee, error) {
	rows, err := repo.db.Query("SELECT * FROM Employee")
	if err != nil {
		return nil, fmt.Errorf("failed to query all employees: %v", err)
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
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

func (repo *employeeRepository) GetByID(id int) (*model.Employee, error) {
	var employee model.Employee
	row := repo.db.QueryRow("SELECT * FROM Employee WHERE ID = :1", id)
	if err := row.Scan(&employee.Id, &employee.Ime, &employee.Prezime, &employee.Email, &employee.DatRodj,
		&employee.Lozinka, &employee.Uloga, &employee.UloZap, &employee.MbrZap); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &employee, nil
}
