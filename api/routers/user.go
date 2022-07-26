package routers

import (
	"clean-architecture/api/controller"
	"clean-architecture/lib"
)

// UserRoutes struct
type UserRoutes struct {
	handler        lib.RequestHandler
	userController controller.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetOneUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	handler lib.RequestHandler,
	userController controller.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
