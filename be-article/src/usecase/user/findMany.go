package userusecase

import (
	"github.com/dionarya23/be-article/src/entities"
)

type PaginatedUser struct {
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
	Total int              `json:"total"`
	Users []*entities.User `json:"users"`
}

func (i *sUserUsecase) FindMany(filters *entities.UserSearchFilter) (interface{}, error) {
	allUsers, err := i.userRepository.FindMany(filters)

	if err != nil {
		return nil, err
	}

	totalCount, err := i.userRepository.Count(filters)
	if err != nil {
		return nil, err
	}

	totalPages := totalCount / filters.Limit
	if totalCount%filters.Limit != 0 {
		totalPages++
	}

	return &PaginatedUser{
		Users: allUsers,
		Page:  filters.Page,
		Limit: filters.Limit,
		Total: totalCount,
	}, nil
}
