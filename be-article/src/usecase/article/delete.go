package articleusecase

import "github.com/dionarya23/be-article/src/entities"

func (i *sArticleUsecase) Delete(articleId *int, authorId *int) error {
	filters := entities.ArticleSearchFilter{
		ID:       *articleId,
		AuthorID: *authorId,
	}

	cat, _ := i.articleRepository.IsExists(&filters)

	if !cat {
		return ErrArticleNotFound
	}

	err := i.articleRepository.SoftDelete(articleId)

	if err != nil {
		return err
	}

	return nil
}
