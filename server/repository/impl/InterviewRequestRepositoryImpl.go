package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type interviewRequestRepository struct {
	db *sql.DB
}

func NewInterviewRequestRepository(db *sql.DB) repository.InterviewRequestRepository {
	return &interviewRequestRepository{db}
}

func (repo *interviewRequestRepository) GetAll() ([]model.InterviewRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_INTERVJU") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&interviewRequest.StatusPozInt, &interviewRequest.RazOdbPozInt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		interviewRequests = append(interviewRequests, interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) GetByID(id int) (*model.InterviewRequest, error) {
	var interviewRequest model.InterviewRequest
	row := repo.db.QueryRow("SELECT * FROM POZIV_NA_INTERVJU WHERE IDPOZINT = :1", id)
	if err := row.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
		&interviewRequest.StatusPozInt, &interviewRequest.RazOdbPozInt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &interviewRequest, nil
}

func (repo *interviewRequestRepository) GetAllBySenderID(userID int) ([]model.InterviewRequest, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_INTERVJU") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&interviewRequest.StatusPozInt, &interviewRequest.RazOdbPozInt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		interviewRequests = append(interviewRequests, interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) GetAllByReceiverID(userID int) ([]model.InterviewRequest, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_INTERVJU") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&interviewRequest.StatusPozInt, &interviewRequest.RazOdbPozInt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		interviewRequests = append(interviewRequests, interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) Create(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("INSERT INTO POZIV_NA_INTERVJU (IDPOZINT, MESODRINT, DATVREPOZINT, STATUSPOZINT, RAZODBPOZINT) "+
		"VALUES (:1, :2, :3, :4, :5)", interviewRequest.IdPozInt, interviewRequest.MesOdrPozInt, interviewRequest.DatVrePozInt,
		interviewRequest.StatusPozInt, interviewRequest.RazOdbPozInt)
	if err != nil {
		return fmt.Errorf("failed to create a interview request: %v", err)
	}
	return nil
}

func (repo *interviewRequestRepository) Update(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("UPDATE POZIV_NA_INTERVJU SET MESODRINT = :1, DATVREPOZINT = :2, STATUSPOZINT = :3,"+
		" RAZODBPOZINT = :4 WHERE IDPOZINT = :5", interviewRequest.IdPozInt, interviewRequest.MesOdrPozInt,
		interviewRequest.DatVrePozInt, interviewRequest.StatusPozInt, interviewRequest.RazOdbPozInt)
	if err != nil {
		return fmt.Errorf("failed to update interview request: %v", err)
	}
	return nil
}
