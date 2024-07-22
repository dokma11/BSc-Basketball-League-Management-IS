package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type trainingRepository struct {
	db *sql.DB
}

func NewTrainingRepository(db *sql.DB) repository.TrainingRepository {
	return &trainingRepository{db}
}

func (repo *trainingRepository) GetAll() ([]model.Training, error) {
	rows, err := repo.db.Query("SELECT * FROM TRENING") // proveriti naziv samo
	if err != nil {
		return nil, fmt.Errorf("failed to query all trainings: %v", err)
	}
	defer rows.Close()

	var trainings []model.Training
	for rows.Next() {
		var training model.Training
		if err := rows.Scan(&training.IdTrng, &training.TrajTrng, &training.DatVreTrng, &training.MesOdrTrng,
			&training.BelesTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		trainings = append(trainings, training)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainings, nil
}

func (repo *trainingRepository) GetByID(id int) (*model.Training, error) {
	var training model.Training
	row := repo.db.QueryRow("SELECT * FROM TRENING WHERE IDTRNG = :1", id)
	if err := row.Scan(&training.IdTrng, &training.TrajTrng, &training.DatVreTrng, &training.MesOdrTrng,
		&training.BelesTrng); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &training, nil
}

func (repo *trainingRepository) GetAllByUserID(userID int) ([]model.Training, error) {
	// TODO: Implementirati ovu metodu kada se spoji sve kako treba (za sada je samo kao GetAll())
	rows, err := repo.db.Query("SELECT * FROM TRENING") // ovde treba dodati idTima
	if err != nil {
		return nil, fmt.Errorf("failed to query all trainings: %v", err)
	}
	defer rows.Close()

	var trainings []model.Training
	for rows.Next() {
		var training model.Training
		if err := rows.Scan(&training.IdTrng, &training.TrajTrng, &training.DatVreTrng, &training.MesOdrTrng,
			&training.BelesTrng); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		trainings = append(trainings, training)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return trainings, nil
}

func (repo *trainingRepository) Create(training *model.Training) error {
	_, err := repo.db.Exec("INSERT INTO TRENING (IDTRNG, TRAJTRNG, DATVRETRNG, MESODTRNG, BELESTRNG) "+
		"VALUES (:1, :2, :3, :4, :5)", training.IdTrng, training.TrajTrng, training.DatVreTrng,
		training.MesOdrTrng, training.BelesTrng)
	if err != nil {
		return fmt.Errorf("failed to create a training: %v", err)
	}
	return nil
}
