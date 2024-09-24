package articlerepository

import "github.com/dionarya23/be-article/src/entities"

func (i *sArticleRepository) Count(filters *entities.ArticleSearchFilter) (int, error) {
	query := "SELECT COUNT(*) FROM articles WHERE deleted_at IS NULL"
	params := []interface{}{}

	if filters.ID != 0 {
		query += " AND id = ?"
		params = append(params, filters.ID)
	}
	if filters.Search != "" {
		query += " AND title LIKE ?"
		params = append(params, "%"+filters.Search+"%")
	}

	var count int
	err := i.DB.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
