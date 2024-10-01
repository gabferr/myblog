package models

import "time"

type User struct {
	ID        int64      `db:"id"`
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	Email     string     `db:email`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"` // Modificado para aceitar NULL
}
