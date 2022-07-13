package repository

import (
	"context"
	"database/sql"
	"posts-parser/internal/jsonlog"
	"posts-parser/internal/models"
	"time"
)

type PostRepository struct {
	db      *sql.DB
	timeout time.Duration
	logger  *jsonlog.Logger
}

func newPostRepo(db *sql.DB, timeout time.Duration, logger *jsonlog.Logger) *PostRepository {
	return &PostRepository{
		db:      db,
		timeout: timeout,
		logger:  logger,
	}
}

func (u *PostRepository) CreatePosts(posts []*models.Post) error {

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return models.ErrDBError
	}
	batch, err := tx.Prepare("INSERT INTO posts (post)")
	if err != nil {
		return models.ErrDBError
	}

	for i := 0; i < len(posts); i++ {
		_, err = batch.Exec(
			posts[i],
		)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				return models.ErrDBError
			}
			return models.ErrDBError
		}
	}
	err = tx.Commit()
	if err != nil {
		return models.ErrDBError
	}
	return nil
}
