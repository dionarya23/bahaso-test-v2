package userv1controller

import (
	"net/http"
	"strconv"

	"github.com/dionarya23/be-article/src/entities"
	userrepository "github.com/dionarya23/be-article/src/repositories/user"

	userusecase "github.com/dionarya23/be-article/src/usecase/user"
	"github.com/labstack/echo/v4"
)

func (i *V1User) FindMany(c echo.Context) (err error) {
	filters := &entities.UserSearchFilter{}
	filters.Limit = 10
	filters.Page = 1

	if pageStr := c.QueryParam("page"); pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Page = page
	}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}

	if search := c.QueryParam("search"); search != "" {
		filters.Search = search
	}

	uu := userusecase.New(
		userrepository.New(i.DB),
	)

	data, err := uu.FindMany(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "User found successfully",
		Data:    data,
	})
}
