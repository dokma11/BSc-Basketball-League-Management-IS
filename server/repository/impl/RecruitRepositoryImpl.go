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
	position := fromRecruitPosition(recruit)
	_, err := repo.db.Exec("INSERT INTO REGRUT (ID, KONTELEFON, "+
		"VISREG, TEZREG, POZREG, PROSRANKREG, PROSOCREG, IDDRAFT) "+
		"VALUES (:1, :2, :3, :4, :5, :6, :7, :8)", recruit.ID, recruit.PhoneNumber, recruit.Height, recruit.Weight,
		position, recruit.AverageRank, recruit.AverageGrade, recruit.DraftId)
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

func (repo *recruitRepository) AddToWishlist(recruit *model.Recruit, teamId int) error {
	_, err := repo.db.Exec(`INSERT INTO ZELJATIMA VALUES (0, SYSDATE, 'REGRUT', 3, NULL, NULL, NULL, :1, :2)`, teamId, recruit.ID) // 3 is recruit type
	if err != nil {
		return fmt.Errorf("failed to add recruit to the wishlist: %v", err)
	}
	return nil
}

func (repo *recruitRepository) RemoveFromWishlist(recruit *model.Recruit, teamId int) error {
	_, err := repo.db.Exec(`DELETE FROM ZELJATIMA WHERE IDREGRUT = :1 AND IDTIM = :2`, recruit.ID, teamId)
	if err != nil {
		return fmt.Errorf("failed to remove recruit from the wishlist: %v", err)
	}
	return nil
}

func (repo *recruitRepository) GetAllByName(name string) ([]model.Recruit, error) {
	rows, err := repo.db.Query(`SELECT DISTINCT K.ID, K.IME, K.PREZIME, K.EMAIL, K.DATRODJ, K.LOZINKA, K.ULOGA,
									R.KONTELEFON, R.VISREG, R.TEZREG, R.POZREG, R.PROSRANKREG, R.PROSOCREG, R.IDDRAFT
									FROM KORISNIK K
									JOIN REGRUT R ON K.ID = R.ID
									WHERE UPPER(K.IME) LIKE UPPER(:1)
									   OR UPPER(K.PREZIME) LIKE UPPER(:2)
									   OR UPPER(K.IME) LIKE UPPER('%' || :3 || '%')
									   OR UPPER(K.PREZIME) LIKE UPPER('%' || :4 || '%')
									   OR UPPER(K.IME || ' ' || K.PREZIME) LIKE UPPER(:5)
								       OR UPPER(K.IME || ' ' || K.PREZIME) LIKE UPPER('%' || :6 || '%')`, name, name, name, name, name, name)
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

func fromRecruitPosition(recruitDAO *model.Recruit) string {
	var position string
	switch recruitDAO.Position {
	case 0:
		position = "PG"
	case 1:
		position = "SG"
	case 2:
		position = "SF"
	case 3:
		position = "PF"
	default:
		position = "C"
	}
	return position
}
