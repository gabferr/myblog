package db

import (
    "database/sql"
    "fmt"
    "time"
    _ "github.com/glebarez/go-sqlite"
)

func Initialize() (*sql.DB, *sql.DB, error) {
    // Lógica de inicialização do banco de dados, similar ao código fornecido anteriormente
}

func UpdateTimestamps(db *sql.DB, tableName string, id int64) error {
    // Implementação similar ao código fornecido anteriormente
}

func SoftDelete(db *sql.DB, tableName string, id int64) error {
    // Implementação similar ao código fornecido anteriormente
}""