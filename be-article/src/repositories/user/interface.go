package userrepository

import (
	"database/sql"
	"time"

	"github.com/dionarya23/be-article/src/entities"
)

type sUserRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(*ParamsCreateUser) (*entities.User, error)
	FindOne(*entities.ParamsCreateUser) (*entities.User, error)
	IsExists(*entities.ParamsCreateUser) (bool, error)
	UpdateResetToken(userID int64, token string, expiry time.Time) error
	UpdatePassword(userId int64, password string) error
}

func New(db *sql.DB) UserRepository {
	return &sUserRepository{DB: db}
}
