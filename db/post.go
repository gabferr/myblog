package db

import (
	"database/sql"
	"time"

	"github.com/gabferr/myblog/models"
)

func CreatePost(db *sql.DB, post *models.Post) error {
	query := `
        INSERT INTO posts (user_id, title, content, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)
    `
	now := time.Now()
	result, err := db.Exec(query, post.UserID, post.Title, post.Content, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = id
	post.CreatedAt = now
	post.UpdatedAt = now

	return nil
}

func UpdatePost(db *sql.DB, post *models.Post) error {
	query := `
        UPDATE posts
        SET title = ?, content = ?, updated_at = ?
        WHERE id = ? AND deleted_at IS NULL
    `
	now := time.Now()
	_, err := db.Exec(query, post.Title, post.Content, now, post.ID)
	if err != nil {
		return err
	}

	post.UpdatedAt = now

	return nil
}

func DeletePost(db *sql.DB, postID int64) error {
	query := `
        UPDATE posts
        SET deleted_at = ?
        WHERE id = ? AND deleted_at IS NULL
    `
	now := time.Now()
	_, err := db.Exec(query, now, postID)
	return err
}

// Função adicional para obter um post por ID
func GetPostByID(db *sql.DB, id int64) (*models.Post, error) {
	query := "SELECT id, user_id, title, content, created_at, updated_at, deleted_at FROM posts WHERE id = ?"
	post := &models.Post{}
	err := db.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetAllPosts busca todos os posts no banco de dados
func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	query := "SELECT id, user_id, title, content, created_at, updated_at, deleted_at FROM posts WHERE deleted_at IS NULL ORDER BY created_at DESC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
