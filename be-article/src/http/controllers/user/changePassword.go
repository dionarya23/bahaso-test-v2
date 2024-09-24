package userv1controller

import (
	"net/http"

	userUsecase "github.com/dionarya23/be-article/src/usecase/user"

	userRepository "github.com/dionarya23/be-article/src/repositories/user"
	"github.com/labstack/echo/v4"
)

func (i *V1User) ChangePassword(c echo.Context) (err error) {
	u := new(changePassword)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uu := userUsecase.New(
		userRepository.New(i.DB),
	)

	data, err := uu.ChangePassword(u.Password, u.Token)

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Password change successfully",
		Data:    data,
	})
}

type (
	changePassword struct {
		Password string `json:"password" validate:"required"`
		Token    string `json:"token" validate:"required"`
	}
)
