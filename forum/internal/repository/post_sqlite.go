package repository

import (
	"database/sql"
	"fmt"

	"lincoln.boris/forum/models"
)

type PostSQLite struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostSQLite {
	return &PostSQLite{db: db}
}

func (r *PostSQLite) Create(post models.Post) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, body, created) VALUES (?, ?, CURRENT_TIMESTAMP)", postTable)

	result, err := tx.Exec(query, post.Title, post.Body)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
/*
	query = fmt.Sprintf("INSERT INTO %s (post_id, category_id) VALUES (?, ?)", postCategoryTable)

	for _, category := range post.Categories {
		_, err = tx.Exec(query, postId, category.CategoryId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
*/
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(postId), nil
}

func (r *PostSQLite) GetAll() ([]models.Post, error) {
	var posts []models.Post

	query := fmt.Sprintf("SELECT post_id, title, body, created FROM %s", postTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.Created)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostSQLite) GetById(post_id int) (models.Post, error) {
	var post models.Post

	query := fmt.Sprintf("SELECT post_id, title, body, created FROM %s WHERE post_id = ?", postTable)

	row := r.db.QueryRow(query, post_id)

	err := row.Scan(&post.Id, &post.Title, &post.Body, &post.Created)
	if err == sql.ErrNoRows {
		return models.Post{}, models.ErrNoRecord
	} else if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *PostSQLite) Update(post_id int, post models.Post) error {
	query := fmt.Sprintf("UPDATE %s SET title = ?, body = ? WHERE post_id = ?", postTable)

	// Delete existing rows in post_category table.
	// Add new rows into post_category table.

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, post.Title, post.Body, post_id)
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

func (r *PostSQLite) Delete(post_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = ?", postTable)

	tx, err := r.db.Begin()
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

/*
func (p *PostSQLite) GetCategoryById(post_id int) ([]models.PostCategory, error) {
	query := `
	SELECT post_category.category_id, category.name
	FROM post_category
	INNER JOIN category ON post_category.category_id = category.category_id
	WHERE post_category.post_id = ?`

	rows, err := p.db.Query(query, post_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var post_categories []models.PostCategory

	for rows.Next() {
		var pc models.PostCategory

		err := rows.Scan(&pc.CategoryId, &pc.CategoryName)
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		} else if err != nil {
			return nil, err
		}

		post_categories = append(post_categories, pc)
	}

	return post_categories, nil
}
*/
