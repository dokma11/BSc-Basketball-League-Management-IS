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
	rows, err := repo.db.Query("SELECT * FROM PozivNaIntervju WHERE IDREGRUT = :1", userID)
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
	_, err := repo.db.Exec(`INSERT INTO PozivNaIntervju (IDPOZINT, MESODRPOZINT, DATVREPOZINT, STATUSPOZINT, RAZODBPOZINT, IDREGRUT, IDTRENER)
		VALUES (0, :1, :2, 'WAITING', NULL, :3, :4)`, interviewRequest.OccurrenceLocation, interviewRequest.OccurrenceDateTime,
		interviewRequest.RecruitId, interviewRequest.CoachId)
	if err != nil {
		return fmt.Errorf("failed to create a interview request: %v", err)
	}
	return nil
}

func (repo *interviewRequestRepository) Update(interviewRequest *model.InterviewRequest) error {
	status := fromStatusEnum(interviewRequest)
	_, err := repo.db.Exec("UPDATE PozivNaIntervju SET STATUSPOZINT = :1, RAZODBPOZINT = :2 WHERE IDPOZINT = :3",
		status, &interviewRequest.DenialReason, interviewRequest.ID)
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

func fromStatusEnum(interviewRequest *model.InterviewRequest) string {
	var status string
	switch interviewRequest.Status {
	case 0:
		status = "WAITING"
	case 1:
		status = "AFFIRMED"
	default:
		status = "REJECTED"
	}
	return status
}
