package articlev1controller

import (
	"net/http"
	"strconv"

	articlerepository "github.com/dionarya23/be-article/src/repositories/article"
	userrepository "github.com/dionarya23/be-article/src/repositories/user"

	articleusecase "github.com/dionarya23/be-article/src/usecase/article"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func (i *V1Article) Delete(c echo.Context) (err error) {
	idArticle := c.Param("id")
	id, err := strconv.Atoi(idArticle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'id'",
		})
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	uu := articleusecase.New(
		articlerepository.New(i.DB),
		userrepository.New(i.DB),
	)

	err = uu.Delete(&id, &uid.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Article deleted successfully",
		Data:    nil,
	})
}
