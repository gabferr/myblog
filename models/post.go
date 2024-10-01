package models

import "time"

type Post struct {
	ID        int64      `db:"id"`
	UserID    int64      `db:"user_id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"` // Modificado para aceitar NULL
}
