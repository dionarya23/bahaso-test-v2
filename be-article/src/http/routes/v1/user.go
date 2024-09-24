package v1routes

import (
	userv1controller "github.com/dionarya23/be-article/src/http/controllers/user"
	"github.com/dionarya23/be-article/src/http/middlewares"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userController := userv1controller.New(&userv1controller.V1User{
		DB: i.DB,
	})

	g.GET("", userController.FindMany, middlewares.Authentication([]string{"author", "visitor"}))
	g.POST("/register", userController.Register)
	g.POST("/login", userController.Login)
	g.POST("/forgot-password", userController.ForgotPassword)
	g.POST("/change-password", userController.ChangePassword)
}
