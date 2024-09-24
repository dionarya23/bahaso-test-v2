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

func (i *V1Article) Update(c echo.Context) (err error) {
	u := new(createRequest)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'id'",
		})
	}

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

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	uu := articleusecase.New(
		articlerepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.Update(&id, &entities.ParamsCreateArticle{
		Title:    u.Title,
		Content:  u.Content,
		ImageURL: u.ImageUrl,
		AuthorID: uid.ID,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Article updated successfully",
		Data:    data,
	})
}
