package articleusecase

import (
	"github.com/dionarya23/be-article/src/entities"
)

type PaginatedArticles struct {
	Page     int                 `json:"page"`
	Limit    int                 `json:"limit"`
	Total    int                 `json:"total"`
	Articles []*entities.Article `json:"articles"`
}

func (i *sArticleUsecase) FindMany(filters *entities.ArticleSearchFilter) (interface{}, error) {
	allArticles, err := i.articleRepository.FindMany(filters)

	if err != nil {
		return nil, err
	}

	totalCount, err := i.articleRepository.Count(filters)
	if err != nil {
		return nil, err
	}

	totalPages := totalCount / filters.Limit
	if totalCount%filters.Limit != 0 {
		totalPages++
	}

	var response interface{}
	response = &PaginatedArticles{
		Articles: allArticles,
		Page:     filters.Page,
		Limit:    filters.Limit,
		Total:    totalCount,
	}
	if filters.ID != 0 {
		response = allArticles[0]
	}

	return response, nil
}
