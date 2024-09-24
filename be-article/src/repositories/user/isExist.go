package userrepository

import (
	"log"
	"strings"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sUserRepository) IsExists(filters *entities.ParamsCreateUser) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE "
	params := []interface{}{}
	conditions := []string{}

	if filters.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, filters.ID)
	}

	if filters.Email != "" {
		conditions = append(conditions, "email = ?")
		params = append(params, filters.Email)
	}

	if len(conditions) > 0 {
		query += strings.Join(conditions, " AND ")
	} else {
		query += "1=1"
	}
	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if user exists: %s", err)
		return false, err
	}

	return exists, nil
}
