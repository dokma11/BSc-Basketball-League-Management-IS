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
	rows, err := repo.db.Query("SELECT * FROM PozivNaIntervju")
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		var status string
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&status, &interviewRequest.RazOdbPozInt, &interviewRequest.IdRegrut, &interviewRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "WAITING" {
			interviewRequest.StatusPozInt = 0
		} else if status == "AFFIRMED" {
			interviewRequest.StatusPozInt = 1
		} else if status == "REJECTED" {
			interviewRequest.StatusPozInt = 2
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
	var status string
	row := repo.db.QueryRow("SELECT * FROM PozivNaIntervju WHERE IDPOZINT = :1", id)
	if err := row.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
		&status, &interviewRequest.RazOdbPozInt, &interviewRequest.IdRegrut, &interviewRequest.IdTrener); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	if status == "WAITING" {
		interviewRequest.StatusPozInt = 0
	} else if status == "AFFIRMED" {
		interviewRequest.StatusPozInt = 1
	} else if status == "REJECTED" {
		interviewRequest.StatusPozInt = 2
	}

	return &interviewRequest, nil
}

func (repo *interviewRequestRepository) GetAllBySenderID(userID int) ([]model.InterviewRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaIntervju WHERE IDTRENER = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		var status string
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&status, &interviewRequest.RazOdbPozInt, &interviewRequest.IdRegrut, &interviewRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "WAITING" {
			interviewRequest.StatusPozInt = 0
		} else if status == "AFFIRMED" {
			interviewRequest.StatusPozInt = 1
		} else if status == "REJECTED" {
			interviewRequest.StatusPozInt = 2
		}

		interviewRequests = append(interviewRequests, interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) GetAllByReceiverID(userID int) ([]model.InterviewRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaIntervju = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequest model.InterviewRequest
		var status string
		if err := rows.Scan(&interviewRequest.IdPozInt, &interviewRequest.MesOdrPozInt, &interviewRequest.DatVrePozInt,
			&status, &interviewRequest.RazOdbPozInt, &interviewRequest.IdRegrut, &interviewRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "WAITING" {
			interviewRequest.StatusPozInt = 0
		} else if status == "AFFIRMED" {
			interviewRequest.StatusPozInt = 1
		} else if status == "REJECTED" {
			interviewRequest.StatusPozInt = 2
		}

		interviewRequests = append(interviewRequests, interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) Create(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("INSERT INTO PozivNaIntervju (IDPOZINT, MESODRINT, DATVREPOZINT, STATUSPOZINT, RAZODBPOZINT, "+
		"IDREGRUT, IDTRENER) "+
		"VALUES (:1, :2, :3, :4, :5)", interviewRequest.IdPozInt, interviewRequest.MesOdrPozInt, interviewRequest.DatVrePozInt,
		interviewRequest.StatusPozInt, &interviewRequest.RazOdbPozInt, interviewRequest.IdRegrut, interviewRequest.IdTrener)
	if err != nil {
		return fmt.Errorf("failed to create a interview request: %v", err)
	}
	return nil
}

func (repo *interviewRequestRepository) Update(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("UPDATE PozivNaIntervju SET MESODRINT = :2, DATVREPOZINT = :3, STATUSPOZINT = :4,"+
		" RAZODBPOZINT = :5 WHERE IDPOZINT = :1", interviewRequest.IdPozInt, interviewRequest.MesOdrPozInt,
		interviewRequest.DatVrePozInt, interviewRequest.StatusPozInt, interviewRequest.RazOdbPozInt)
	if err != nil {
		return fmt.Errorf("failed to update interview request: %v", err)
	}
	return nil
}
