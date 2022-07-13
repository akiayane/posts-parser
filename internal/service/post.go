package service

import (
	"posts-parser/config"
	"posts-parser/internal/jsonlog"
	"posts-parser/internal/models"
	"posts-parser/internal/repository"
)

type PostService struct {
	postRepo repository.Post
	cfg      *config.Configs
	logger   *jsonlog.Logger
}

func (m PostService) CreatePosts(posts []*models.Post) error {
	return m.postRepo.CreatePosts(posts)
}

func NewPostService(postRepo repository.Post, cfg *config.Configs, logger *jsonlog.Logger) *PostService {
	return &PostService{
		postRepo: postRepo,
		cfg:      cfg,
		logger:   logger,
	}
}
