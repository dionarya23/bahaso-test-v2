package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"

	"github.com/dionarya23/be-article/src/helpers"
)

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func Authentication(approvedRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

			if token == "" {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			claims, err := helpers.ValidateJWT(&helpers.ParamsValidateJWT{
				Token:     token,
				SecretKey: os.Getenv("JWT_SECRET"),
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			user := make(map[string]interface{})
			mapstructure.Decode(claims, &user)

			role, ok := user["role"].(string)
			if !ok {
				return c.JSON(http.StatusForbidden, ErrorResponse{
					Status:  false,
					Message: "Invalid role format",
				})
			}

			if len(approvedRoles) > 0 {
				if !slices.Contains(approvedRoles, role) {
					return errors.New("You don't have an access!")
				}
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
