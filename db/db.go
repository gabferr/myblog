package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

var dbConn *sql.DB

// SetDB inicializa a conexão de banco de dados para os handlers
func SetDB(conn *sql.DB) {
	dbConn = conn
}

func Initialize() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./blog.db")
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	// Defina a conexão como a conexão global
	SetDB(db)

	return db, nil
}

func createTables(db *sql.DB) error {
	// Criação da tabela de usuários
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP
    )`)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela de usuários: %v", err)
	}

	// Criação da tabela de posts
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    )`)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela de posts: %v", err)
	}

	return nil
}

func UpdateTimestamps(tableName string, id int64) error {
	now := time.Now().Format(time.RFC3339)
	query := fmt.Sprintf(`UPDATE %s SET updated_at = ? WHERE id = ?`, tableName)
	_, err := dbConn.Exec(query, now, id)
	return err
}

func SoftDelete(tableName string, id int64) error {
	now := time.Now().Format(time.RFC3339)
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = ? WHERE id = ?`, tableName)
	_, err := dbConn.Exec(query, now, id)
	return err
}
