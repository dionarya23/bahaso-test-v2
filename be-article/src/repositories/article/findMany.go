package articlerepository

import (
	"log"
	"reflect"
	"strings"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sArticleRepository) FindMany(filters *entities.ArticleSearchFilter) ([]*entities.Article, error) {
	query := "SELECT id, author_id, title, content, image_url, created_at, updated_at FROM articles WHERE deleted_at IS NULL and 1=1"
	params := []interface{}{}

	n := &entities.ArticleSearchFilter{}

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.ID != 0 {
			conditions = append(conditions, "id = ?")
			params = append(params, filters.ID)
		}
		if filters.Search != "" {
			conditions = append(conditions, "name LIKE ?")
			params = append(params, "%"+filters.Search+"%")
		}

		if len(conditions) > 0 {
			query += " AND " + strings.Join(conditions, " AND ")
		}
	}

	query += " ORDER BY created_at DESC"

	offset := (filters.Page - 1) * filters.Limit
	query += " LIMIT ? OFFSET ?"
	params = append(params, filters.Limit, offset)

	if filters.Offset != 0 {
		query += " OFFSET ?"
		params = append(params, filters.Offset)
	}

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding article: %s", err)
		return nil, err
	}
	defer rows.Close()

	articles := make([]*entities.Article, 0)
	for rows.Next() {
		article := new(entities.Article)
		err := rows.Scan(&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.ImageURL, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		article.IsOwned = int(article.AuthorID) == filters.AuthorID
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
