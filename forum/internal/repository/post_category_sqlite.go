package repository

import (
	"database/sql"
	"fmt"

	"lincoln.boris/forum/models"
)

type PostCategorySQLite struct {
	db *sql.DB
}

func NewPostCategorySQLite(db *sql.DB) *PostCategorySQLite {
	return &PostCategorySQLite{db: db}
}


func (pc *PostCategorySQLite) Create(post_id, category_id int) (int, error) {
	tx, err := pc.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (post_id, category_id, created) VALUES (?, ?, CURRENT_TIMESTAMP)", postCategoryTable)

	result, err := tx.Exec(query, post_id, category_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	postCategoryId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(postCategoryId), nil
}

func (pc *PostCategorySQLite) GetAll(post_id int) ([]models.PostCategory, error) {
	var post_categories []models.PostCategory

	query := fmt.Sprintf("SELECT post_id, category_id FROM %s WHERE post_id = ?", postCategoryTable)

	rows, err := pc.db.Query(query, post_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p models.PostCategory

		err := rows.Scan(&p.PostId, &p.CategoryId)
		if err != nil {
			return nil, err
		}

		post_categories = append(post_categories, p)
	}

	return post_categories, nil
}

func (pc *PostCategorySQLite) Delete(post_id, category_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = ? AND category_id = ?", postCategoryTable)

	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, post_id, category_id)
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

func (pc *PostCategorySQLite) DeleteAll(post_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = ?", postCategoryTable)

	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, post_id)
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
