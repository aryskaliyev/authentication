package repository

import (
	"database/sql"

	"lincoln.boris/forum/models"
)

type Post interface {
	Create(post models.Post) (int, error)
	GetAll() ([]models.Post, error)
	GetById(postId int) (models.Post, error)
	Update(postId int, post models.Post) error
	Delete(postId int) error
//	GetCategoryById(postId int) ([]models.PostCategory, error)
}

type Category interface {
	Create(category models.Category) (int, error)
	GetAll() ([]models.Category, error)
	GetById(categoryId int) (models.Category, error)
	Update(categoryId int, category models.Category) error
	Delete(categoryId int) error
}

type PostCategory interface {
	Create(postId, categoryId int) (int, error)
	GetAll(postId int) ([]models.PostCategory, error)
	Delete(postId, categoryId int) error
	DeleteAll(postId int) error
}

type Repository struct {
	Post
	Category
	PostCategory
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:     NewPostSQLite(db),
		Category: NewCategorySQLite(db),
		PostCategory: NewPostCategorySQLite(db),
	}
}
