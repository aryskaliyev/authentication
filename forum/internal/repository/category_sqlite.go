package repository

import (
	"database/sql"
	"fmt"

	"lincoln.boris/forum/models"
)

type CategorySQLite struct {
	db *sql.DB
}

func NewCategorySQLite(db *sql.DB) *CategorySQLite {
	return &CategorySQLite{db: db}
}

func (c *CategorySQLite) Create(input models.Category) (int, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (name, created) VALUES (?, CURRENT_TIMESTAMP)", categoryTable)

	result, err := tx.Exec(query, input.Name)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	categoryId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(categoryId), nil
}

func (c *CategorySQLite) GetAll() ([]models.Category, error) {
	var categories []models.Category

	query := fmt.Sprintf("SELECT category_id, name, created FROM %s", categoryTable)

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Category

		err := rows.Scan(&category.Id, &category.Name, &category.Created)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategorySQLite) GetById(category_id int) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT category_id, name, created FROM %s WHERE category_id = ?", categoryTable)

	row := c.db.QueryRow(query, category_id)

	err := row.Scan(&category.Id, &category.Name, &category.Created)
	if err == sql.ErrNoRows {
		return models.Category{}, models.ErrNoRecord
	} else if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (c *CategorySQLite) Update(category_id int, category models.Category) error {
	query := fmt.Sprintf("UPDATE %s SET name = ? WHERE category_id = ?", categoryTable)

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, category.Name, category_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *CategorySQLite) Delete(category_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE category_id = ?", categoryTable)

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, category_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
