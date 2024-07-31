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
	rows, err := repo.db.Query(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, Z.ULOZAP, Z.MBRZAP, Z.IDUGO 
									  FROM ZAPOSLENI Z, KORISNIK K
									  WHERE K.ID = Z.ID`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all employees: %v", err)
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employeeDAO model.EmployeeDAO
		var role, employeeRole string
		if err := rows.Scan(&employeeDAO.ID, &employeeDAO.FirstName, &employeeDAO.LastName, &employeeDAO.Email, &employeeDAO.DateOfBirth,
			&employeeDAO.Password, &role, &employeeRole, &employeeDAO.MbrZap, &employeeDAO.IdUgo); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRoleStrings(role, employeeRole, &employeeDAO)
		employee := &model.Employee{}
		employee.FromDAO(&employeeDAO)

		employees = append(employees, *employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return employees, nil
}

func (repo *employeeRepository) GetByID(id int) (*model.Employee, error) {
	var employeeDAO model.EmployeeDAO
	var role, employeeRole string
	row := repo.db.QueryRow(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, Z.ULOZAP, Z.MBRZAP, Z.IDUGO 
								   FROM ZAPOSLENI Z, KORISNIK K
								   WHERE K.ID = Z.ID AND K.ID = :1`, id)
	if err := row.Scan(&employeeDAO.ID, &employeeDAO.FirstName, &employeeDAO.LastName, &employeeDAO.Email, &employeeDAO.DateOfBirth,
		&employeeDAO.Password, &role, &employeeRole, &employeeDAO.MbrZap, &employeeDAO.IdUgo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRoleStrings(role, employeeRole, &employeeDAO)
	employee := &model.Employee{}
	employee.FromDAO(&employeeDAO)

	return employee, nil
}

func (repo *employeeRepository) GetByTeamID(teamID int) (*model.Employee, error) {
	var employeeDAO model.EmployeeDAO
	var role, employeeRole string
	row := repo.db.QueryRow(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, Z.ULOZAP, Z.MBRZAP, Z.IDUGO
								   FROM ZAPOSLENI Z, UGOVOR U, KORISNIK K
								   WHERE Z.ID = K.ID AND Z.IDUGO = U.IDUGO AND U.IDTIM = :1 AND Z.ULOZAP = 'Menadzer'`, teamID)
	if err := row.Scan(&employeeDAO.ID, &employeeDAO.FirstName, &employeeDAO.LastName, &employeeDAO.Email, &employeeDAO.DateOfBirth,
		&employeeDAO.Password, &role, &employeeRole, &employeeDAO.MbrZap, &employeeDAO.IdUgo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRoleStrings(role, employeeRole, &employeeDAO)
	employee := &model.Employee{}
	employee.FromDAO(&employeeDAO)

	return employee, nil
}

func fromRoleStrings(role string, employeeRole string, employee *model.EmployeeDAO) {
	if role == "Zaposleni" {
		employee.Role = 1
	} else {
		fmt.Println("Error: employee must have an employee role!")
	}

	if employeeRole == "Menadzer" {
		employee.UloZap = 0
	} else if employeeRole == "Igrac" {
		employee.UloZap = 1
	} else if employeeRole == "Trener" {
		employee.UloZap = 2
	} else if employeeRole == "Skaut" {
		employee.UloZap = 3
	}
}
