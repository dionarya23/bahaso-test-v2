package articlerepository

import (
	"log"
	"time"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sArticleRepository) Update(articleId *int, p *entities.ParamsCreateArticle) (*entities.CreateArticle, error) {
	_, err := i.DB.Exec("UPDATE articles SET title = ?, content = ?, image_url = ?, author_id = ? WHERE id = ?",
		p.Title,
		p.Content,
		p.ImageURL,
		p.AuthorID,
		articleId,
	)

	if err != nil {
		log.Printf("Error updating article: %s", err)
		return nil, err
	}

	var id int64
	var createdAt time.Time
	err = i.DB.QueryRow("SELECT id, created_at FROM articles WHERE id = ?", articleId).Scan(&id, &createdAt)
	if err != nil {
		log.Printf("Error retrieving updated article data: %s", err)
		return nil, err
	}

	article := &entities.CreateArticle{
		ID:        id,
		CreatedAt: createdAt,
	}

	return article, nil
}
