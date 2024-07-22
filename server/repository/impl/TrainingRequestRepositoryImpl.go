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
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_TRENING") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&trainingRequest.StatusPozTrng, &trainingRequest.RazOdbPozTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
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
	row := repo.db.QueryRow("SELECT * FROM POZIV_NA_TRENING WHERE IDPOZTRNG = :1", id)
	if err := row.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
		&trainingRequest.StatusPozTrng, &trainingRequest.RazOdbPozTrng); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &trainingRequest, nil
}

func (repo *trainingRequestRepository) GetAllBySenderID(userID int) ([]model.TrainingRequest, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_TRENING") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&trainingRequest.StatusPozTrng, &trainingRequest.RazOdbPozTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
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
	rows, err := repo.db.Query("SELECT * FROM POZIV_NA_TRENING") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all training requests: %v", err)
	}
	defer rows.Close()

	var trainingRequests []model.TrainingRequest
	for rows.Next() {
		var trainingRequest model.TrainingRequest
		if err := rows.Scan(&trainingRequest.IdPozTrng, &trainingRequest.DatVrePozTrng, &trainingRequest.MesOdrPozTrng,
			&trainingRequest.StatusPozTrng, &trainingRequest.RazOdbPozTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		trainingRequests = append(trainingRequests, trainingRequest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainingRequests, nil
}

func (repo *trainingRequestRepository) Create(trainingRequest *model.TrainingRequest) error {
	_, err := repo.db.Exec("INSERT INTO POZIV_NA_TRENING (IDPOZTRNG, DATVREPOZTRNG, MESODRPOZTRNG, STATUSPOZTRNG, RAZODBPOZTRNG) "+
		"VALUES (:1, :2, :3, :4, :5)", trainingRequest.IdPozTrng, trainingRequest.DatVrePozTrng, trainingRequest.MesOdrPozTrng,
		trainingRequest.StatusPozTrng, trainingRequest.RazOdbPozTrng)
	if err != nil {
		return fmt.Errorf("failed to create a TrainingRequest: %v", err)
	}
	return nil
}

func (r *trainingRequestRepository) Update(trainingRequest *model.TrainingRequest) error {
	_, err := r.db.Exec("UPDATE POZIV_NA_TRENING SET DATVREPOZTRNG = :1, MESODRPOZTRNG = :2, STATUSPOZTRNG = :3,"+
		" RAZODBPOZTRNG = :4 WHERE IDPOZTRNG = :5", trainingRequest.DatVrePozTrng, trainingRequest.MesOdrPozTrng,
		trainingRequest.StatusPozTrng, trainingRequest.RazOdbPozTrng, trainingRequest.IdPozTrng)
	if err != nil {
		return fmt.Errorf("failed to update tim: %v", err)
	}
	return nil
}
