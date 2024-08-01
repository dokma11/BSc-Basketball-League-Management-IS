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
		var interviewRequestDAO model.InterviewRequestDAO
		var status string
		if err := rows.Scan(&interviewRequestDAO.IdPozInt, &interviewRequestDAO.MesOdrPozInt, &interviewRequestDAO.DatVrePozInt,
			&status, &interviewRequestDAO.RazOdbPozInt, &interviewRequestDAO.IdRegrut, &interviewRequestDAO.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusString(status, &interviewRequestDAO)
		interviewRequest := &model.InterviewRequest{}
		interviewRequest.FromDAO(&interviewRequestDAO)

		interviewRequests = append(interviewRequests, *interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) GetByID(id int) (*model.InterviewRequest, error) {
	var interviewRequestDAO model.InterviewRequestDAO
	var status string
	row := repo.db.QueryRow("SELECT * FROM PozivNaIntervju WHERE IDPOZINT = :1", id)
	if err := row.Scan(&interviewRequestDAO.IdPozInt, &interviewRequestDAO.MesOdrPozInt, &interviewRequestDAO.DatVrePozInt,
		&status, &interviewRequestDAO.RazOdbPozInt, &interviewRequestDAO.IdRegrut, &interviewRequestDAO.IdTrener); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromStatusString(status, &interviewRequestDAO)
	interviewRequest := &model.InterviewRequest{}
	interviewRequest.FromDAO(&interviewRequestDAO)

	return interviewRequest, nil
}

func (repo *interviewRequestRepository) GetAllBySenderID(userID int) ([]model.InterviewRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaIntervju WHERE IDTRENER = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all interview requests: %v", err)
	}
	defer rows.Close()

	var interviewRequests []model.InterviewRequest
	for rows.Next() {
		var interviewRequestDAO model.InterviewRequestDAO
		var status string
		if err := rows.Scan(&interviewRequestDAO.IdPozInt, &interviewRequestDAO.MesOdrPozInt, &interviewRequestDAO.DatVrePozInt,
			&status, &interviewRequestDAO.RazOdbPozInt, &interviewRequestDAO.IdRegrut, &interviewRequestDAO.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusString(status, &interviewRequestDAO)
		interviewRequest := &model.InterviewRequest{}
		interviewRequest.FromDAO(&interviewRequestDAO)

		interviewRequests = append(interviewRequests, *interviewRequest)
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
		var interviewRequestDAO model.InterviewRequestDAO
		var status string
		if err := rows.Scan(&interviewRequestDAO.IdPozInt, &interviewRequestDAO.MesOdrPozInt, &interviewRequestDAO.DatVrePozInt,
			&status, &interviewRequestDAO.RazOdbPozInt, &interviewRequestDAO.IdRegrut, &interviewRequestDAO.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromStatusString(status, &interviewRequestDAO)
		interviewRequest := &model.InterviewRequest{}
		interviewRequest.FromDAO(&interviewRequestDAO)

		interviewRequests = append(interviewRequests, *interviewRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return interviewRequests, nil
}

func (repo *interviewRequestRepository) Create(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("INSERT INTO PozivNaIntervju (IDPOZINT, MESODRINT, DATVREPOZINT, STATUSPOZINT, RAZODBPOZINT, "+
		"IDREGRUT, IDTRENER) "+
		"VALUES (:1, :2, :3, :4, :5)", interviewRequest.ID, interviewRequest.OccurrenceLocation, interviewRequest.OccurrenceDateTime,
		interviewRequest.Status, &interviewRequest.DenialReason, interviewRequest.RecruitId, interviewRequest.CoachId)
	if err != nil {
		return fmt.Errorf("failed to create a interview request: %v", err)
	}
	return nil
}

func (repo *interviewRequestRepository) Update(interviewRequest *model.InterviewRequest) error {
	_, err := repo.db.Exec("UPDATE PozivNaIntervju SET MESODRINT = :2, DATVREPOZINT = :3, STATUSPOZINT = :4,"+
		" RAZODBPOZINT = :5 WHERE IDPOZINT = :1", interviewRequest.ID, interviewRequest.OccurrenceLocation,
		interviewRequest.OccurrenceDateTime, interviewRequest.Status, &interviewRequest.DenialReason)
	if err != nil {
		return fmt.Errorf("failed to update interview request: %v", err)
	}
	return nil
}

func fromStatusString(status string, interviewRequest *model.InterviewRequestDAO) {
	switch status {
	case "WAITING":
		interviewRequest.StatusPozInt = 0
	case "AFFIRMED":
		interviewRequest.StatusPozInt = 1
	default:
		interviewRequest.StatusPozInt = 2
	}
}
