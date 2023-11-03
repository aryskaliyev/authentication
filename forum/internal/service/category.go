package service

import (
	"lincoln.boris/forum/internal/repository"
	"lincoln.boris/forum/models"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) Create(category models.Category) (int, error) {
	return c.repo.Create(category)
}

func (c *CategoryService) GetAll() ([]models.Category, error) {
	return c.repo.GetAll()
}

func (c *CategoryService) GetById(category_id int) (models.Category, error) {
	return c.repo.GetById(category_id)
}

func (c *CategoryService) Update(category_id int, category models.Category) error {
	_, err := c.GetById(category_id)
	if err == models.ErrNoRecord {
		return models.ErrNoRecord
	}

	return c.repo.Update(category_id, category)
}

func (c *CategoryService) Delete(category_id int) error {
	_, err := c.GetById(category_id)
	if err == models.ErrNoRecord {
		return models.ErrNoRecord
	}

	return c.repo.Delete(category_id)
}
