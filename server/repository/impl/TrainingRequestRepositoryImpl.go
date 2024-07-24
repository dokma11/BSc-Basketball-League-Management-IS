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
	rows, err := repo.db.Query("SELECT * FROM PozivNaTrening") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		var status string
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&status, &trainingRequest.RazOdbPozTrng, &trainingRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "PENDING" {
			trainingRequest.StatusPozTrng = 0
		} else if status == "APPROVED" {
			trainingRequest.StatusPozTrng = 1
		} else if status == "DISAPPROVED" {
			trainingRequest.StatusPozTrng = 2
		}

		trainingRequests = append(trainingRequests, trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) GetByID(id int) (*model.TrainingRequest, error) {
	var trainingRequest model.TrainingRequest
	var status string
	row := repo.db.QueryRow("SELECT * FROM PozivNaTrening WHERE IDPOZTRNG = :1", id)
	if err := row.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
		&status, &trainingRequest.RazOdbPozTrng, &trainingRequest.IdTrener); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	if status == "PENDING" {
		trainingRequest.StatusPozTrng = 0
	} else if status == "APPROVED" {
		trainingRequest.StatusPozTrng = 1
	} else if status == "DISAPPROVED" {
		trainingRequest.StatusPozTrng = 2
	}

	return &trainingRequest, nil
}

func (repo *trainingRequestRepository) GetAllBySenderID(userID int) ([]model.TrainingRequest, error) {
	rows, err := repo.db.Query("SELECT * FROM PozivNaTrening WHERE IDTRENER = :1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		var status string
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&status, &trainingRequest.RazOdbPozTrng, &trainingRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "PENDING" {
			trainingRequest.StatusPozTrng = 0
		} else if status == "APPROVED" {
			trainingRequest.StatusPozTrng = 1
		} else if status == "DISAPPROVED" {
			trainingRequest.StatusPozTrng = 2
		}

		trainingRequests = append(trainingRequests, trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) GetAllByReceiverID(userID int) ([]model.TrainingRequest, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM PozivNaTrening") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		var status string
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&status, &trainingRequest.RazOdbPozTrng, &trainingRequest.IdTrener); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if status == "PENDING" {
			trainingRequest.StatusPozTrng = 0
		} else if status == "APPROVED" {
			trainingRequest.StatusPozTrng = 1
		} else if status == "DISAPPROVED" {
			trainingRequest.StatusPozTrng = 2
		}

		trainingRequests = append(trainingRequests, trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) Create(trainingRequest *model.TrainingRequest) error {
	_, err := repo.db.Exec("INSERT INTO PozivNaTrening (IDPOZTRNG, DATVREPOZTRNG, MESODRPOZTRNG, STATUSPOZTRNG, RAZODBPOZTRNG, "+
		"IDTRENER) VALUES (:1, :2, :3, :4, :5, :6)", trainingRequest.IdPozTrng, trainingRequest.DatVrePozTrng,
		trainingRequest.MesOdrPozTrng, trainingRequest.StatusPozTrng, &trainingRequest.RazOdbPozTrng, &trainingRequest.IdTrener)
	if err != nil {
		return fmt.Errorf("failed to create a training request: %v", err)
	}
	return nil
}

func (repo *trainingRequestRepository) Update(trainingRequest *model.TrainingRequest) error {
	_, err := repo.db.Exec("UPDATE PozivNaTrening SET DATVREPOZTRNG = :1, MESODRPOZTRNG = :2, STATUSPOZTRNG = :3,"+
		" RAZODBPOZTRNG = :4 WHERE IDPOZTRNG = :5", trainingRequest.DatVrePozTrng, trainingRequest.MesOdrPozTrng,
		trainingRequest.StatusPozTrng, trainingRequest.RazOdbPozTrng, trainingRequest.IdPozTrng)
	if err != nil {
		return fmt.Errorf("failed to update training request: %v", err)
	}
	return nil
}
