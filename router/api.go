package router

import (
	"ecommerce-backend/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo			*echo.Echo
	UserHandler		handler.UserHandler
}

func (api *API) SetupRouter() {
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
}
