package v1routes

import (
	articlev1controller "github.com/dionarya23/be-article/src/http/controllers/article"
	"github.com/dionarya23/be-article/src/http/middlewares"
)

func (i *V1Routes) MountArticle() {
	g := i.Echo.Group("/article")

	articleController := articlev1controller.New(&articlev1controller.V1Article{
		DB: i.DB,
	})
	g.GET("", articleController.GetArticle)
	g.GET("/admin", articleController.GetArticle, middlewares.Authentication([]string{"author", "visitor"}))
	g.POST("", articleController.Create, middlewares.Authentication([]string{"author"}))
	g.PUT("/:id", articleController.Update, middlewares.Authentication([]string{"author"}))
	g.DELETE("/:id", articleController.Delete, middlewares.Authentication([]string{"author"}))
}
