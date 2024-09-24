package articlerepository

import (
	"database/sql"

	"github.com/dionarya23/be-article/src/entities"
)

type sArticleRepository struct {
	DB *sql.DB
}

type ArticleRepository interface {
	Create(*entities.ParamsCreateArticle) (*entities.CreateArticle, error)
	FindMany(*entities.ArticleSearchFilter) ([]*entities.Article, error)
	IsExists(*entities.ArticleSearchFilter) (bool, error)
	Update(*int, *entities.ParamsCreateArticle) (*entities.CreateArticle, error)
	Count(filters *entities.ArticleSearchFilter) (int, error)
	SoftDelete(*int) error
}

func New(db *sql.DB) ArticleRepository {
	return &sArticleRepository{DB: db}
}
