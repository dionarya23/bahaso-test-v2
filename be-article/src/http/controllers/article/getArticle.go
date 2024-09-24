package articlev1controller

import (
	"net/http"
	"strconv"

	"github.com/dionarya23/be-article/src/entities"
	articlerepository "github.com/dionarya23/be-article/src/repositories/article"
	userrepository "github.com/dionarya23/be-article/src/repositories/user"

	articleusecase "github.com/dionarya23/be-article/src/usecase/article"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type (
	meValidator struct {
		ID int `mapstructure:"user_id" validate:"required"`
	}
)

func (i *V1Article) GetArticle(c echo.Context) (err error) {
	filters := &entities.ArticleSearchFilter{}
	filters.Limit = 10
	filters.Page = 1
	if idStr := c.QueryParam("id"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'id'",
			})
		}
		filters.ID = id
	}

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
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)
	filters.AuthorID = uid.ID

	if search := c.QueryParam("search"); search != "" {
		filters.Search = search
	}

	uu := articleusecase.New(
		articlerepository.New(i.DB),
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
		Message: "Article found successfully",
		Data:    data,
	})
}
