package articlerepository

import (
	"log"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sArticleRepository) IsExists(filters *entities.ArticleSearchFilter) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM articles WHERE 1=1 "
	params := []interface{}{}

	if filters.ID != 0 {
		query += "AND id = ? "
		params = append(params, filters.ID)
	}
	if filters.Search != "" {
		query += "AND title = ? "
		params = append(params, filters.Search)
	}
	if filters.AuthorID != 0 {
		query += "AND author_id = ? "
		params = append(params, filters.AuthorID)
	}

	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if article exists: %s", err)
		return false, err
	}

	return exists, nil
}
