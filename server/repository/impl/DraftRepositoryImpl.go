package impl

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"database/sql"
	"errors"
	"fmt"
)

type draftRepository struct {
	db *sql.DB
}

func NewDraftRepository(db *sql.DB) repository.DraftRepository {
	return &draftRepository{db}
}

func (repo *draftRepository) GetAll() ([]model.Draft, error) {
	rows, err := repo.db.Query("SELECT * FROM DRAFT")
	if err != nil {
		return nil, fmt.Errorf("failed to query all drafts: %v", err)
	}
	defer rows.Close()

	var drafts []model.Draft
	for rows.Next() {
		var draft model.Draft
		if err := rows.Scan(&draft.IdDraft, &draft.GodOdrDraft, &draft.LokOdrDraft); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		drafts = append(drafts, draft)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return drafts, nil
}

func (repo *draftRepository) GetByID(id int) (*model.Draft, error) {
	var draft model.Draft
	row := repo.db.QueryRow("SELECT * FROM DRAFT WHERE IDDRAFT = :1", id)
	if err := row.Scan(&draft.IdDraft, &draft.GodOdrDraft, &draft.LokOdrDraft); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No result found
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}
	return &draft, nil
}