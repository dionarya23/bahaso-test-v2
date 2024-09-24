package userrepository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/dionarya23/be-article/src/entities"
)

func (i *sUserRepository) FindOne(filters *entities.ParamsCreateUser) (*entities.User, error) {
	query := "SELECT id, name, email, password, role FROM users WHERE "
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

	if filters.Token != "" {
		conditions = append(conditions, "reset_token = ?")
		params = append(params, filters.Token)
	}

	if len(conditions) > 0 {
		query += strings.Join(conditions, " AND ")
	} else {
		query += "1"
	}

	query += " LIMIT 1"

	row := i.DB.QueryRow(query, params...)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		log.Printf("Error find user: %s", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
