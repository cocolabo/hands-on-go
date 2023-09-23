package router

import (
	"github.com/cocolabo/hands-on-go/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SineUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)

	return e
}
