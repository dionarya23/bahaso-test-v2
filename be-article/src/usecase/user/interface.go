package userusecase

import (
	"github.com/dionarya23/be-article/src/entities"
	user "github.com/dionarya23/be-article/src/repositories/user"
)

type sUserUsecase struct {
	userRepository user.UserRepository
}

type UserUsecase interface {
	CreateUser(*ParamsCreateUser) (*ResultLogin, error)
	Login(*ParamsLogin) (*ResultLogin, error)
	ForgotPassword(string) (bool, error)
	ChangePassword(string, string) (bool, error)
	FindMany(*entities.UserSearchFilter) (interface{}, error)
}

func New(
	userRepository user.UserRepository,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
