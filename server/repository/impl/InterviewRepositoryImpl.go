package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type interviewRepository struct {
	db *sql.DB
}

func NewInterviewRepository(db *sql.DB) repository.InterviewRepository {
	return &interviewRepository{db}
}

func (repo *interviewRepository) GetAll() ([]model.Interview, error) {
	rows, err := repo.db.Query("SELECT * FROM INTERVJU") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all interviews: %v", err)
	}
	defer rows.Close()

	var interviews []model.Interview
	for rows.Next() {
		var interview model.Interview
		if err := rows.Scan(&interview.IdInt, &interview.MesOdrInt, &interview.DatVreInt, &interview.BelesInt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		interviews = append(interviews, interview)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviews, nil
}

func (repo *interviewRepository) GetByID(id int) (*model.Interview, error) {
	var interview model.Interview
	row := repo.db.QueryRow("SELECT * FROM INTERVJU WHERE IDINT = :1", id)
	if err := row.Scan(&interview.IdInt, &interview.MesOdrInt, &interview.DatVreInt, &interview.BelesInt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &interview, nil
}

func (repo *interviewRepository) GetAllByUserID(userID int) ([]model.Interview, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM INTERVJU") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all interviews: %v", err)
	}
	defer rows.Close()

	var interviews []model.Interview
	for rows.Next() {
		var interview model.Interview
		if err := rows.Scan(&interview.IdInt, &interview.MesOdrInt, &interview.DatVreInt, &interview.BelesInt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		interviews = append(interviews, interview)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviews, nil
}

func (repo *interviewRepository) Create(interview *model.Interview) error {
	_, err := repo.db.Exec("INSERT INTO INTERVJU (IDINT, MESODRINT, DATVREINT, BELESINT) "+
		"VALUES (:1, :2, :3, :4, :5)", interview.IdInt, interview.MesOdrInt, interview.DatVreInt, interview.BelesInt)
	if err != nil {
		return fmt.Errorf("failed to create a interview: %v", err)
	}
	return nil
}
