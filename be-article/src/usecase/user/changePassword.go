package userusecase

import (
	"github.com/dionarya23/be-article/src/entities"
	"github.com/dionarya23/be-article/src/helpers"
)

func (i *sUserUsecase) ChangePassword(password string, token string) (bool, error) {
	filters := entities.ParamsCreateUser{
		Token: token,
	}

	user, _ := i.userRepository.FindOne(&filters)
	if user == nil {
		return false, ErrUserNotFound
	}

	hashedPassword, _ := helpers.HashPassword(password)
	err := i.userRepository.UpdatePassword(user.ID, hashedPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}
