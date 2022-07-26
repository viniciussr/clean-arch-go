package api

import (
	"clean-architecture/api/controller"
	"clean-architecture/api/routers"
	"clean-architecture/infrastructure/repository"
	"clean-architecture/pkg"
	user "clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {

	userRepo := repository.NewUserInMemory()
	userService := user.NewService(userRepo)

	//user
	controller := controller.NewUserController(*userService)

	handler := gin.New()

	userRouter := routers.NewUserRoutes(handler, controller)
	userRouter.Setup()

	httpServer := pkg.New(handler)

	// Shutdown
	httpServer.Shutdown()

}
