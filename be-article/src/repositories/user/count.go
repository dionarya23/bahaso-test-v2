package userrepository

import "github.com/dionarya23/be-article/src/entities"

func (i *sUserRepository) Count(filters *entities.UserSearchFilter) (int, error) {
	query := "SELECT COUNT(*) FROM users WHERE 1=1"
	params := []interface{}{}

	if filters.Search != "" {
		query += " AND name LIKE ?"
		params = append(params, "%"+filters.Search+"%")
	}

	var count int
	err := i.DB.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
