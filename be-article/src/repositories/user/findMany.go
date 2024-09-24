package userrepository

import (
	"log"
	"reflect"
	"strings"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sUserRepository) FindMany(filters *entities.UserSearchFilter) ([]*entities.User, error) {
	query := "SELECT id, name, email, role, created_at, updated_at FROM users WHERE 1=1"
	params := []interface{}{}

	n := &entities.UserSearchFilter{}

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

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
		log.Printf("Error finding user: %s", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)
	for rows.Next() {
		user := new(entities.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
