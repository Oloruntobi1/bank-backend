// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Account struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
	// must be positive
	Balance   int64     `json:"balance"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}