package main

import (
	"database/sql"
	"fmt"
	"log"

	"clean-architecture/api/controller"
	"clean-architecture/api/routers"
	"clean-architecture/infrastructure/repository"
	"clean-architecture/lib"
	user "clean-architecture/usecase"
)

func main() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", "clean_architecture_go", "clean_architecture_go", "127.0.0.1", "clean_architecture_go")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	userRepo := repository.NewUserMySQL(db)
	userService := user.NewService(userRepo)

	//user
	controller := controller.NewUserController(*userService)

	handler := lib.NewRequestHandler()

	userRouter := routers.NewUserRoutes(lib.NewRequestHandler(), controller)
	userRouter.Setup()

	handler.Gin.Run()
}
