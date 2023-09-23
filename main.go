package main

import (
	"github.com/cocolabo/hands-on-go/controller"
	"github.com/cocolabo/hands-on-go/db"
	"github.com/cocolabo/hands-on-go/repository"
	"github.com/cocolabo/hands-on-go/router"
	"github.com/cocolabo/hands-on-go/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
