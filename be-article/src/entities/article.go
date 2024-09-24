package entities

import (
	"time"
)

type Article struct {
	ID        int64     `json:"id" db:"id"`
	AuthorID  int64     `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	ImageURL  string    `json:"image_url,omitempty" db:"image_url"`
	IsOwned   bool      `json:"is_owned"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ArticleSearchFilter struct {
	ID       int
	Page     int
	Limit    int
	Offset   int
	AuthorID int
	Search   string
}

type ParamsCreateArticle struct {
	Title    string
	Content  string
	ImageURL string
	AuthorID int
}

type CreateArticle struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
