package db

import (
	"database/sql"
	"time"

	"github.com/gabferr/myblog/models"
)

// GetUserByUsername busca um usuário no banco de dados pelo nome de usuário
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, password, email, created_at FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuário não encontrado
		}
		return nil, err
	}
	return user, nil
}

// CreateUser insere um novo usuário no banco de dados
func CreateUser(db *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)"
	now := time.Now()
	result, err := db.Exec(query, user.Username, user.Email, user.Password, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

// GetPostsByUserID busca todos os posts de um usuário com base no ID do usuário
func GetPostsByUserID(db *sql.DB, userID int64) ([]*models.Post, error) {
	query := `
        SELECT id, title, content, created_at, updated_at
        FROM posts
        WHERE user_id = ? AND deleted_at IS NULL
        ORDER BY created_at DESC
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
