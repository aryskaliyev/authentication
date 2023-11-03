package service

import (
	"lincoln.boris/forum/internal/repository"
	"lincoln.boris/forum/models"
)

type Post interface {
	Create(input models.Post) (int, error)
	GetAll() ([]models.Post, error)
	GetById(postId int) (models.Post, error)
	Update(postId int, input models.Post) error
	Delete(postId int) error
}

type Category interface {
	Create(input models.Category) (int, error)
	GetAll() ([]models.Category, error)
	GetById(categoryId int) (models.Category, error)
	Update(categoryId int, input models.Category) error
	Delete(categoryId int) error
}

type PostCategory interface {
	Create(postId, categoryId int) (int, error)
	GetAll(postId int) ([]models.PostCategory, error)
	Delete(postId, categoryId int) error
	DeleteAll(postId int) error
}

type Service struct {
	Post
	Category
	PostCategory
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post:         NewPostService(repos.Post),
		Category:     NewCategoryService(repos.Category),
		PostCategory: NewPostCategoryService(repos.PostCategory),
	}
}
