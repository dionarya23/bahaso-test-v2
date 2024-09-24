package articlerepository

import (
	"log"
	"time"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sArticleRepository) Create(p *entities.ParamsCreateArticle) (*entities.CreateArticle, error) {
	res, err := i.DB.Exec("INSERT INTO articles (author_id, title, content, image_url) VALUES (?, ?, ?, ?)",
		p.AuthorID,
		p.Title,
		p.Content,
		p.ImageURL,
	)

	if err != nil {
		log.Printf("Error creating article: %s", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %s", err)
		return nil, err
	}

	var createdAt time.Time
	err = i.DB.QueryRow("SELECT created_at FROM articles WHERE id = ?", id).Scan(&createdAt)
	if err != nil {
		log.Printf("Error retrieving created_at: %s", err)
		return nil, err
	}

	article := &entities.CreateArticle{
		ID:        id,
		CreatedAt: createdAt,
	}

	return article, nil
}
