package entities

import "time"

type User struct {
	ID                int64     `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	Email             string    `json:"email" db:"email"`
	Password          string    `json:"-" db:"password"`
	Role              string    `json:"role" db:"role"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	ResetToken        string    `json:"-" db:"reset_token"`
	ResetTokenExpires time.Time `json:"-" db:"reset_token_expires"`
}

type ParamsCreateUser struct {
	ID    int64
	Name  string
	Email string
	Token string
}

type UserSearchFilter struct {
	Page   int
	Limit  int
	Offset int
	Search string
}
