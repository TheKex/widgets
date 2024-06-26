// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Counter struct {
	ID        int64         `json:"id"`
	WidgetID  int64         `json:"widget_id"`
	Order     sql.NullInt32 `json:"order"`
	Value     sql.NullInt32 `json:"value"`
	Increment sql.NullInt32 `json:"increment"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Widget struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
