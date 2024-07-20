package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type regrutRepository struct {
	db *sql.DB
}

func NewRegrutRepository(db *sql.DB) repository.RegrutRepository {
	return &regrutRepository{db}
}

func (repo *regrutRepository) GetAll() ([]model.Regrut, error) {
	rows, err := repo.db.Query("SELECT * FROM REGRUT")
	if err != nil {
		return nil, fmt.Errorf("failed to query all recruits: %v", err)
	}
	defer rows.Close()

	var recruits []model.Regrut
	for rows.Next() {
		var recruit model.Regrut
		if err := rows.Scan(&recruit.Id, &recruit.Ime, &recruit.Prezime, &recruit.Email, &recruit.DatRodj,
			&recruit.Lozinka, &recruit.Uloga, &recruit.KonTelefonReg, &recruit.MesRodjReg, &recruit.VisReg,
			&recruit.TezReg, &recruit.PozReg, &recruit.ProsRankReg, &recruit.ProsOcReg); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		recruits = append(recruits, recruit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return recruits, nil
}

func (repo *regrutRepository) GetByID(id int) (*model.Regrut, error) {
	var recruit model.Regrut
	row := repo.db.QueryRow("SELECT * FROM REGRUT WHERE ID = :1", id)
	if err := row.Scan(&recruit.Id, &recruit.Ime, &recruit.Prezime, &recruit.Email, &recruit.DatRodj,
		&recruit.Lozinka, &recruit.Uloga, &recruit.KonTelefonReg, &recruit.MesRodjReg, &recruit.VisReg,
		&recruit.TezReg, &recruit.PozReg, &recruit.ProsRankReg, &recruit.ProsOcReg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &recruit, nil
}

func (repo *regrutRepository) Create(recruit *model.Regrut) error {
	_, err := repo.db.Exec("INSERT INTO REGRUT (ID, EMAIL, IME, PREZIME, DATRODJ, LOZINKA, ULOGA, KONTELEFONREG, "+
		"MESRODJREG, VISREG, TEZREG, POZREG, PROSRANKREG, PROSOCREG) "+
		"VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13)", recruit.Id, recruit.Ime, recruit.Prezime,
		recruit.Email, recruit.DatRodj, recruit.Lozinka, recruit.Uloga, recruit.KonTelefonReg, recruit.MesRodjReg,
		recruit.VisReg, recruit.TezReg, recruit.PozReg, recruit.ProsRankReg, recruit.ProsOcReg)
	if err != nil {
		return fmt.Errorf("failed to create recruit: %v", err)
	}
	return nil
}

func (repo *regrutRepository) Update(recruit *model.Regrut) error {
	_, err := repo.db.Exec("UPDATE REGRUT SET IME = :2, PREZIME = :3, EMAIL = :4, DATRODJ = :5, LOZINKA = :6,"+
		"ULOGA = :7,  KONTELEFONREG = :8, MESRODJREG = :9, VISREG = :10, TEZREG = :11, POZREG = :12, PROSRANKREG = :13, "+
		"PROSOCREG = :14 WHERE ID = :1", recruit.Id, recruit.Ime, recruit.Prezime, recruit.Email, recruit.DatRodj,
		recruit.Lozinka, recruit.Uloga, recruit.KonTelefonReg, recruit.MesRodjReg, recruit.VisReg, recruit.TezReg,
		recruit.PozReg, recruit.ProsRankReg, recruit.ProsOcReg)
	if err != nil {
		return fmt.Errorf("failed to update recruit: %v", err)
	}
	return nil
}
