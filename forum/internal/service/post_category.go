package service

import (
	"lincoln.boris/forum/internal/repository"
	"lincoln.boris/forum/models"
)

type PostCategoryService struct {
	repo repository.PostCategory
}

func NewPostCategoryService(repo repository.PostCategory) *PostCategoryService {
	return &PostCategoryService{repo: repo}
}

func (s *PostCategoryService) Create(post_id, category_id int) (int, error) {}

func (s *PostCategoryService) GetAll(post_id int) ([]models.PostCategory, error) {}

func (s *PostCategoryService) Delete(post_id, category_id int) error {}

func (s *PostCategoryService) DeleteAll(post_id int) error {}
