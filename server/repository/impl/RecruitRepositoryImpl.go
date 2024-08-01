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
	rows, err := repo.db.Query("SELECT * FROM REGRUT")
	if err != nil {
		return nil, fmt.Errorf("failed to query all recruits: %v", err)
	}
	defer rows.Close()

	var recruits []model.Recruit
	for rows.Next() {
		var recruitDAO model.RecruitDAO
		var position string
		if err := rows.Scan(&recruitDAO.ID, &recruitDAO.KonTelefonReg, &recruitDAO.VisReg, &recruitDAO.TezReg, &position,
			&recruitDAO.ProsRankReg, &recruitDAO.ProsOcReg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRecruitPositionString(position, &recruitDAO)
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
	var position string
	row := repo.db.QueryRow("SELECT * FROM REGRUT WHERE ID = :1", id)
	if err := row.Scan(&recruitDAO.ID, &recruitDAO.KonTelefonReg, &recruitDAO.VisReg, &recruitDAO.TezReg, &position,
		&recruitDAO.ProsRankReg, &recruitDAO.ProsOcReg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRecruitPositionString(position, &recruitDAO)
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

func fromRecruitPositionString(position string, recruitDAO *model.RecruitDAO) {
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
}
