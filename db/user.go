package db

import (
    "database/sql"
    "github.com/gabferr/myblog/models"
    "time"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
    // Implementação para obter um usuário pelo nome de usuário
}

func CreateUser(db *sql.DB, user *models.User) error {
    // Implementação para criar um novo usuário
}

func GetPostsByUserID(db *sql.DB, userID int64) ([]*models.Post, error) {
    // Implementação para obter todos os posts de um usuário
}

func CreatePost(db *sql.DB, post *models.Post) error {
    // Implementação para criar um novo post
}