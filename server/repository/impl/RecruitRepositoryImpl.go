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
	rows, err := repo.db.Query("SELECT * FROM Recruit")
	if err != nil {
		return nil, fmt.Errorf("failed to query all recruits: %v", err)
	}
	defer rows.Close()

	var recruits []model.Recruit
	for rows.Next() {
		var recruit model.Recruit
		if err := rows.Scan(&recruit.Id, &recruit.Ime, &recruit.Prezime, &recruit.Email, &recruit.DatRodj,
			&recruit.Lozinka, &recruit.Uloga, &recruit.KonTelefonReg, &recruit.VisReg, &recruit.TezReg, &recruit.PozReg,
			&recruit.ProsRankReg, &recruit.ProsOcReg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		recruits = append(recruits, recruit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return recruits, nil
}

func (repo *recruitRepository) GetByID(id int) (*model.Recruit, error) {
	var recruit model.Recruit
	row := repo.db.QueryRow("SELECT * FROM Recruit WHERE ID = :1", id)
	if err := row.Scan(&recruit.Id, &recruit.Ime, &recruit.Prezime, &recruit.Email, &recruit.DatRodj,
		&recruit.Lozinka, &recruit.Uloga, &recruit.KonTelefonReg, &recruit.VisReg, &recruit.TezReg, &recruit.PozReg,
		&recruit.ProsRankReg, &recruit.ProsOcReg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &recruit, nil
}

func (repo *recruitRepository) Create(recruit *model.Recruit) error {
	_, err := repo.db.Exec("INSERT INTO Recruit (ID, EMAIL, IME, PREZIME, DATRODJ, LOZINKA, ULOGA, KONTELEFONREG, "+
		"MESRODJREG, VISREG, TEZREG, POZREG, PROSRANKREG, PROSOCREG) "+
		"VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13)", recruit.Id, recruit.Ime, recruit.Prezime,
		recruit.Email, recruit.DatRodj, recruit.Lozinka, recruit.Uloga, recruit.KonTelefonReg, recruit.VisReg,
		recruit.TezReg, recruit.PozReg, recruit.ProsRankReg, recruit.ProsOcReg)
	if err != nil {
		return fmt.Errorf("failed to create recruit: %v", err)
	}
	return nil
}

func (repo *recruitRepository) Update(recruit *model.Recruit) error {
	_, err := repo.db.Exec("UPDATE Recruit SET IME = :2, PREZIME = :3, EMAIL = :4, DATRODJ = :5, LOZINKA = :6,"+
		"ULOGA = :7,  KONTELEFONREG = :8, VISREG = :9, TEZREG = :10, POZREG = :11, PROSRANKREG = :12, "+
		"PROSOCREG = :13 WHERE ID = :1", recruit.Id, recruit.Ime, recruit.Prezime, recruit.Email, recruit.DatRodj,
		recruit.Lozinka, recruit.Uloga, recruit.KonTelefonReg, recruit.VisReg, recruit.TezReg, recruit.PozReg,
		recruit.ProsRankReg, recruit.ProsOcReg)
	if err != nil {
		return fmt.Errorf("failed to update recruit: %v", err)
	}
	return nil
}
