package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"docker/app/config"
	_userController "docker/user/delivery/http"
	mid "docker/user/delivery/http/middleware"
	"docker/user/repository"
	"docker/user/usecase"
	"os"
)

func Run(){

	db := config.InitDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	e := echo.New()
	mid.NewGoMiddleware().LogMiddleware(e)
	_userController.NewUserController(e, userUsecase)
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	if err := e.Start(address); err != nil {
		log.Info("Exit The Server")
	}
}
