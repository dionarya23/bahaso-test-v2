package articlev1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Article struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Article interface {
	Create(c echo.Context) error
	GetArticle(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

func New(v1Article *V1Article) iV1Article {
	return v1Article
}
