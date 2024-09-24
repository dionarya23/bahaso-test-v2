package articleusecase

import (
	"github.com/dionarya23/be-article/src/entities"
	article "github.com/dionarya23/be-article/src/repositories/article"
	user "github.com/dionarya23/be-article/src/repositories/user"
)

type sArticleUsecase struct {
	articleRepository article.ArticleRepository
	userRepository    user.UserRepository
}

type ArticleUsecase interface {
	Create(*entities.ParamsCreateArticle) (*ResultCreate, error)
	FindMany(*entities.ArticleSearchFilter) (interface{}, error)
	Update(*int, *entities.ParamsCreateArticle) (*ResultCreate, error)
	Delete(*int, *int) error
}

func New(articleRepository article.ArticleRepository, userRepository user.UserRepository) ArticleUsecase {
	return &sArticleUsecase{
		articleRepository: articleRepository,
		userRepository:    userRepository,
	}
}
