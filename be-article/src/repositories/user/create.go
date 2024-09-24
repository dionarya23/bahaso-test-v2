package userrepository

import (
	"log"

	"github.com/dionarya23/be-article/src/entities"
)

type (
	ParamsCreateUser struct {
		Name     string
		Email    string
		Password string
		Role     string
	}
)

func (i *sUserRepository) Create(p *ParamsCreateUser) (*entities.User, error) {
	result, err := i.DB.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)", p.Name, p.Email, p.Password, p.Role)
	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %s", err)
		return nil, err
	}

	user := &entities.User{
		ID:    id,
		Name:  p.Name,
		Email: p.Email,
		Role:  p.Role,
	}

	return user, nil
}
