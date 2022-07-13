package repository

import (
	"database/sql"
	"posts-parser/internal/jsonlog"
	"posts-parser/internal/models"
	"time"
)

type Repository struct {
	Post
}

type Post interface {
	CreatePosts(posts []*models.Post) error
}

func NewRepository(db *sql.DB, timeout time.Duration, logger *jsonlog.Logger) *Repository {
	return &Repository{
		Post: newPostRepo(db, timeout, logger),
	}
}
