package articlev1controller

import (
	"net/http"

	"github.com/dionarya23/be-article/src/entities"
	articlerepository "github.com/dionarya23/be-article/src/repositories/article"
	userrepository "github.com/dionarya23/be-article/src/repositories/user"

	articleusecase "github.com/dionarya23/be-article/src/usecase/article"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func (i *V1Article) Create(c echo.Context) (err error) {
	u := new(createRequest)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uu := articleusecase.New(
		articlerepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.Create(&entities.ParamsCreateArticle{
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

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Article created successfully",
		Data:    data,
	})
}

type (
	createRequest struct {
		Title    string `json:"title" validate:"required"`
		Content  string `json:"content" validate:"required"`
		ImageUrl string `json:"image_url" validate:"required"`
	}
)
