package userusecase

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/dionarya23/be-article/src/entities"
	"github.com/dionarya23/be-article/src/helpers"
)

func generateResetToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (i *sUserUsecase) ForgotPassword(email string) (bool, error) {

	filters := entities.ParamsCreateUser{
		Email: email,
	}

	user, _ := i.userRepository.FindOne(&filters)

	if user == nil {
		return false, ErrUserNotFound
	}

	token, err := generateResetToken()
	if err != nil {
		return false, err
	}
	tokenExpire := time.Now().Add(24 * time.Hour)

	err = i.userRepository.UpdateResetToken(user.ID, token, tokenExpire)
	if err != nil {
		return false, err
	}

	err = helpers.SendMail(user.Email, token)
	if err != nil {
		return false, err
	}

	return true, nil
}
