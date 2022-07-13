package service

import (
	"posts-parser/config"
	"posts-parser/internal/jsonlog"
	"posts-parser/internal/models"
	"posts-parser/internal/repository"
)

type Service struct {
	Post
}

type Post interface {
	CreatePosts(posts []*models.Post) error
}

func NewService(repo *repository.Repository, cfg *config.Configs, logger *jsonlog.Logger) *Service {
	return &Service{
		Post: NewPostService(repo.Post, cfg, logger),
	}
}
