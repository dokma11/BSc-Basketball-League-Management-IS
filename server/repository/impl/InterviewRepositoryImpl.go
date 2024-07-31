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
	rows, err := repo.db.Query("SELECT * FROM INTERVJU")
	if err != nil {
		return nil, fmt.Errorf("failed to query all interviews: %v", err)
	}
	defer rows.Close()

	var interviews []model.Interview
	for rows.Next() {
		var interviewDAO model.InterviewDAO
		if err := rows.Scan(&interviewDAO.IdInt, &interviewDAO.MesOdrInt, &interviewDAO.DatVreInt, &interviewDAO.BelesInt,
			&interviewDAO.IdPozInt, &interviewDAO.IdRegrut); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		interview := &model.Interview{}
		interview.FromDAO(&interviewDAO)

		interviews = append(interviews, *interview)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviews, nil
}

func (repo *interviewRepository) GetByID(id int) (*model.Interview, error) {
	var interviewDAO model.InterviewDAO
	row := repo.db.QueryRow("SELECT * FROM INTERVJU WHERE IDINT = :1", id)
	if err := row.Scan(&interviewDAO.IdInt, &interviewDAO.MesOdrInt, &interviewDAO.DatVreInt, &interviewDAO.BelesInt,
		&interviewDAO.IdPozInt, &interviewDAO.IdRegrut); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	interview := &model.Interview{}
	interview.FromDAO(&interviewDAO)

	return interview, nil
}

func (repo *interviewRepository) GetAllByUserID(userID int) ([]model.Interview, error) {
	rows, err := repo.db.Query("SELECT * FROM INTERVJU WHERE IDREGRUT = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all interviews: %v", err)
	}
	defer rows.Close()

	var interviews []model.Interview
	for rows.Next() {
		var interviewDAO model.InterviewDAO
		if err := rows.Scan(&interviewDAO.IdInt, &interviewDAO.MesOdrInt, &interviewDAO.DatVreInt, &interviewDAO.BelesInt,
			&interviewDAO.IdPozInt, &interviewDAO.IdRegrut); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		interview := &model.Interview{}
		interview.FromDAO(&interviewDAO)

		interviews = append(interviews, *interview)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviews, nil
}

func (repo *interviewRepository) Create(interview *model.Interview) error {
	_, err := repo.db.Exec("INSERT INTO INTERVJU (IDINT, MESODRINT, DATVREINT, BELESINT, IDPOZINT, IDREGRUT) "+
		"VALUES (:1, :2, :3, :4, :5, :6)", interview.ID, interview.OccurrenceLocation, interview.OccurrenceDateTime,
		interview.Notes, interview.InterviewRequestId, interview.RecruitId)
	if err != nil {
		return fmt.Errorf("failed to create a interview: %v", err)
	}
	return nil
}
