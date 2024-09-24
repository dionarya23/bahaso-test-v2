package v1routes

import (
	userv1controller "github.com/dionarya23/be-article/src/http/controllers/user"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userController := userv1controller.New(&userv1controller.V1User{
		DB: i.DB,
	})

	g.POST("/register", userController.Register)
	g.POST("/login", userController.Login)
	g.POST("/forgot-password", userController.ForgotPassword)
	g.POST("/change-password", userController.ChangePassword)
}
