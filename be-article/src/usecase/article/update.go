package articleusecase

import (
	"github.com/dionarya23/be-article/src/entities"
)

func (i *sArticleUsecase) Update(articleId *int, p *entities.ParamsCreateArticle) (*ResultCreate, error) {
	filters := entities.ArticleSearchFilter{
		ID: *articleId,
	}

	cat, _ := i.articleRepository.IsExists(&filters)

	if !cat {
		return nil, ErrArticleNotFound
	}

	data, err := i.articleRepository.Update(articleId,
		&entities.ParamsCreateArticle{
			Title:    p.Title,
			Content:  p.Content,
			ImageURL: p.ImageURL,
			AuthorID: p.AuthorID,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
	}, nil

}
