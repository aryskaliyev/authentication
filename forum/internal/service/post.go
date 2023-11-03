package service

import (
	"lincoln.boris/forum/internal/repository"
	"lincoln.boris/forum/models"
)

type PostService struct {
	repo    repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(post models.Post) (int, error) {
	return s.repo.Create(post)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	posts, err := s.repo.GetAll()
	if err == models.ErrNoRecord {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetById(post_id int) (models.Post, error) {
	post, err := s.repo.GetById(post_id)
	if err == models.ErrNoRecord {
		return models.Post{}, models.ErrNoRecord
	} else if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (s *PostService) Update(post_id int, post models.Post) error {
	_, err := s.GetById(post_id)
	if err == models.ErrNoRecord {
		return models.ErrNoRecord
	}

	return s.repo.Update(post_id, post)
}

func (s *PostService) Delete(post_id int) error {
	_, err := s.GetById(post_id)
	if err == models.ErrNoRecord {
		return models.ErrNoRecord
	}

	return s.repo.Delete(post_id)
}
