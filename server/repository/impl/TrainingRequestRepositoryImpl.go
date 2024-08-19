package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type trainingRequestRepository struct {
	db *sql.DB
}

func NewTrainingRequestRepository(db *sql.DB) repository.TrainingRequestRepository {
	return &trainingRequestRepository{db}
}

func (repo *trainingRequestRepository) GetAll() ([]model.TrainingRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaTrening")
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequestDAO model.TrainingRequestDAO
		var status string
		if err := rows.Scan(&trainingRequestDAO.IdPozTrng, &trainingRequestDAO.DatVrePozTrng, &trainingRequestDAO.MesOdrPozTrng,
			&status, &trainingRequestDAO.RazOdbPozTrng, &trainingRequestDAO.IdTrener, &trainingRequestDAO.TrajTrng, &trainingRequestDAO.NazTipTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRequestStatusString(status, &trainingRequestDAO)
		trainingRequest := &model.TrainingRequest{}
		trainingRequest.FromDAO(&trainingRequestDAO)

		trainingRequests = append(trainingRequests, *trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) GetByID(id int) (*model.TrainingRequest, error) {
	var trainingRequestDAO model.TrainingRequestDAO
	var status string
	row := repo.db.QueryRow("SELECT * FROM PozivNaTrening WHERE IDPOZTRNG = :1", id)
	if err := row.Scan(&trainingRequestDAO.IdPozTrng, &trainingRequestDAO.DatVrePozTrng, &trainingRequestDAO.MesOdrPozTrng,
		&status, &trainingRequestDAO.RazOdbPozTrng, &trainingRequestDAO.IdTrener, &trainingRequestDAO.TrajTrng, &trainingRequestDAO.NazTipTrng); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	fromRequestStatusString(status, &trainingRequestDAO)
	trainingRequest := &model.TrainingRequest{}
	trainingRequest.FromDAO(&trainingRequestDAO)

	return trainingRequest, nil
}

func (repo *trainingRequestRepository) GetAllBySenderID(userID int) ([]model.TrainingRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaTrening WHERE IDTRENER = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequestDAO model.TrainingRequestDAO
		var status string
		if err := rows.Scan(&trainingRequestDAO.IdPozTrng, &trainingRequestDAO.DatVrePozTrng, &trainingRequestDAO.MesOdrPozTrng,
			&status, &trainingRequestDAO.RazOdbPozTrng, &trainingRequestDAO.IdTrener, &trainingRequestDAO.TrajTrng, &trainingRequestDAO.NazTipTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRequestStatusString(status, &trainingRequestDAO)
		trainingRequest := &model.TrainingRequest{}
		trainingRequest.FromDAO(&trainingRequestDAO)

		trainingRequests = append(trainingRequests, *trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) GetAllByReceiverID(userID int) ([]model.TrainingRequest, error) {
	rows, err := repo.db.Query(`SELECT P.IDPOZTRNG, DATVREPOZTRNG, MESODRPOZTRNG, STATUSPOZTRNG, RAZODBPOZTRNG, IDTRENER, TRAJTRNG, NAZTIPTRNG
									  FROM POZIVNATRENING P, upravlja U
									  WHERE P.IDPOZTRNG = U.IDPOZTRNG AND U.IDREGRUT = :1`, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests by receiver id: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequestDAO model.TrainingRequestDAO
		var status string
		if err := rows.Scan(&trainingRequestDAO.IdPozTrng, &trainingRequestDAO.DatVrePozTrng, &trainingRequestDAO.MesOdrPozTrng,
			&status, &trainingRequestDAO.RazOdbPozTrng, &trainingRequestDAO.IdTrener, &trainingRequestDAO.TrajTrng, &trainingRequestDAO.NazTipTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		fromRequestStatusString(status, &trainingRequestDAO)
		trainingRequest := &model.TrainingRequest{}
		trainingRequest.FromDAO(&trainingRequestDAO)

		trainingRequests = append(trainingRequests, *trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) Create(trainingRequest *model.TrainingRequest, recruitId int64) error {
	_, err := repo.db.Exec("INSERT INTO PozivNaTrening (IDPOZTRNG, DATVREPOZTRNG, MESODRPOZTRNG, STATUSPOZTRNG, RAZODBPOZTRNG, "+
		"IDTRENER, TRAJTRNG, NAZTIPTRNG) VALUES (0, :1, :2, 'WAITING', NULL, :3, :4, :5)", trainingRequest.OccurrenceDateTime,
		trainingRequest.OccurrenceLocation, &trainingRequest.CoachId, &trainingRequest.Duration, &trainingRequest.TrainingTypeName)
	if err != nil {
		return fmt.Errorf("failed to create a training request: %v", err)
	}

	// Get the newly created request
	var trainingRequestDAO model.TrainingRequestDAO
	var status string
	row := repo.db.QueryRow("SELECT * FROM PozivNaTrening ORDER BY IDPOZTRNG DESC")
	if err := row.Scan(&trainingRequestDAO.IdPozTrng, &trainingRequestDAO.DatVrePozTrng, &trainingRequestDAO.MesOdrPozTrng,
		&status, &trainingRequestDAO.RazOdbPozTrng, &trainingRequestDAO.IdTrener, &trainingRequestDAO.TrajTrng, &trainingRequestDAO.NazTipTrng); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err // No result found
		}
		return err
	}

	// Create the necessary handling gerund
	_, handlingErr := repo.db.Exec("INSERT INTO UPRAVLJA (IDREGRUT, IDPOZTRNG) VALUES (:1, :2)", recruitId, &trainingRequestDAO.IdPozTrng)
	if handlingErr != nil {
		return fmt.Errorf("failed to create a training request handling gerund: %v", err)
	}
	return nil
}

func (repo *trainingRequestRepository) Update(trainingRequest *model.TrainingRequest) error {
	status := fromTrainingStatusEnum(trainingRequest)
	_, err := repo.db.Exec("UPDATE PozivNaTrening SET STATUSPOZTRNG = :1, RAZODBPOZTRNG = :2 WHERE IDPOZTRNG = :3",
		status, trainingRequest.DenialReason, trainingRequest.ID)
	if err != nil {
		return fmt.Errorf("failed to update training request: %v", err)
	}

	if status == "AFFIRMED" {
		// Get the newly created training
		var trainingId int64
		row := repo.db.QueryRow("SELECT IDTRNG FROM Trening ORDER BY IDTRNG DESC")
		if err := row.Scan(&trainingId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return err // No result found
			}
			return err
		}

		// Get the corresponding training request id
		var recruitId int64
		requestRow := repo.db.QueryRow("SELECT IDREGRUT FROM UPRAVLJA WHERE IDPOZTRNG = :1", trainingRequest.ID)
		if err := requestRow.Scan(&recruitId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return err // No result found
			}
			return err
		}

		// Create the participation gerund
		_, err := repo.db.Exec("INSERT INTO UCESTVUJE VALUES (:1, :2)", recruitId, trainingId)
		if err != nil {
			return fmt.Errorf("failed to insert values into gerund: %v", err)
		}
	}
	return nil
}

func fromRequestStatusString(option string, trainingRequest *model.TrainingRequestDAO) {
	switch option {
	case "WAITING":
		trainingRequest.StatusPozTrng = 0
	case "AFFIRMED":
		trainingRequest.StatusPozTrng = 1
	default:
		trainingRequest.StatusPozTrng = 2
	}
}

func fromTrainingStatusEnum(trainingRequest *model.TrainingRequest) string {
	var status string
	switch trainingRequest.Status {
	case 0:
		status = "WAITING"
	case 1:
		status = "AFFIRMED"
	default:
		status = "REJECTED"
	}
	return status
}
