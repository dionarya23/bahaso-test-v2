package articleusecase

import (
	"time"

	"github.com/dionarya23/be-article/src/entities"
	userusecase "github.com/dionarya23/be-article/src/usecase/user"
)

type (
	ResultCreate struct {
		ID        int64     `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

func (i *sArticleUsecase) Create(p *entities.ParamsCreateArticle) (*ResultCreate, error) {
	filters := entities.ParamsCreateUser{
		ID: int64(p.AuthorID),
	}

	user, _ := i.userRepository.IsExists(&filters)

	if !user {
		return nil, userusecase.ErrInvalidUser
	}

	data, err := i.articleRepository.Create(&entities.ParamsCreateArticle{
		AuthorID: p.AuthorID,
		Title:    p.Title,
		Content:  p.Content,
		ImageURL: p.ImageURL,
	})

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
	}, nil

}
