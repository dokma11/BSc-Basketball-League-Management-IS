package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type recruitRepository struct {
	db *sql.DB
}

func NewRecruitRepository(db *sql.DB) repository.RecruitRepository {
	return &recruitRepository{db}
}

func (repo *recruitRepository) GetAll() ([]model.Recruit, error) {
	rows, err := repo.db.Query(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, R.KONTELEFON,
										R.VISREG, R.TEZREG, R.POZREG, R.PROSRANKREG, R.PROSOCREG, R.IDDRAFT
										FROM KORISNIK K, REGRUT R
										WHERE K.ID = R.ID`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all recruits: %v", err)
	}
	defer rows.Close()

	var recruits []model.Recruit
	for rows.Next() {
		var recruitDAO model.RecruitDAO
		var role, position string
		if err := rows.Scan(&recruitDAO.ID, &recruitDAO.FirstName, &recruitDAO.LastName, &recruitDAO.Email, &recruitDAO.DateOfBirth,
			&recruitDAO.Password, &role, &recruitDAO.KonTelefonReg, &recruitDAO.VisReg, &recruitDAO.TezReg, &position,
			&recruitDAO.ProsRankReg, &recruitDAO.ProsOcReg, &recruitDAO.IdDraft); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRecruitPositionAndRoleString(position, role, &recruitDAO)
		recruit := &model.Recruit{}
		recruit.FromDAO(&recruitDAO)

		recruits = append(recruits, *recruit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return recruits, nil
}

func (repo *recruitRepository) GetByID(id int) (*model.Recruit, error) {
	var recruitDAO model.RecruitDAO
	var role, position string
	row := repo.db.QueryRow(`SELECT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA, R.KONTELEFON,
									R.VISREG, R.TEZREG, R.POZREG, R.PROSRANKREG, R.PROSOCREG, R.IDDRAFT
									FROM KORISNIK K, REGRUT R
									WHERE K.ID = R.ID AND R.ID = :1`, id)
	if err := row.Scan(&recruitDAO.ID, &recruitDAO.FirstName, &recruitDAO.LastName, &recruitDAO.Email, &recruitDAO.DateOfBirth,
		&recruitDAO.Password, &role, &recruitDAO.KonTelefonReg, &recruitDAO.VisReg, &recruitDAO.TezReg, &position,
		&recruitDAO.ProsRankReg, &recruitDAO.ProsOcReg, &recruitDAO.IdDraft); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		fmt.Println(err)
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRecruitPositionAndRoleString(position, role, &recruitDAO)
	recruit := &model.Recruit{}
	recruit.FromDAO(&recruitDAO)

	return recruit, nil
}

func (repo *recruitRepository) Create(recruit *model.Recruit) error {
	_, err := repo.db.Exec("INSERT INTO REGRUT (ID, KONTELEFONREG, "+
		"MESRODJREG, VISREG, TEZREG, POZREG, PROSRANKREG, PROSOCREG) "+
		"VALUES (:1, :2, :3, :4, :5, :6, :7)", recruit.ID, recruit.PhoneNumber, recruit.Height, recruit.Weight,
		recruit.Position, recruit.AverageRank, recruit.AverageGrade)
	if err != nil {
		return fmt.Errorf("failed to create recruit: %v", err)
	}
	return nil
}

func (repo *recruitRepository) Update(recruit *model.Recruit) error {
	_, err := repo.db.Exec("UPDATE Recruit SET KONTELEFONREG = :2, VISREG = :3, TEZREG = :4, POZREG = :5, PROSRANKREG = :6, "+
		"PROSOCREG = :7 WHERE ID = :1", recruit.ID, recruit.PhoneNumber, recruit.Height, recruit.Weight,
		recruit.Position, recruit.AverageRank, recruit.AverageGrade)
	if err != nil {
		return fmt.Errorf("failed to update recruit: %v", err)
	}
	return nil
}

func fromRecruitPositionAndRoleString(position string, role string, recruitDAO *model.RecruitDAO) {
	switch position {
	case "PG":
		recruitDAO.PozReg = 0
	case "SG":
		recruitDAO.PozReg = 1
	case "SF":
		recruitDAO.PozReg = 2
	case "PF":
		recruitDAO.PozReg = 3
	default:
		recruitDAO.PozReg = 4
	}
	switch role {
	case "UloRegrut":
		recruitDAO.Role = 0
	default:
		recruitDAO.Role = 1
	}
}
