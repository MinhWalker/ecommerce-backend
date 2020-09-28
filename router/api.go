package router

import (
	"ecommerce-backend/handler"
	"ecommerce-backend/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo			*echo.Echo
	UserHandler		handler.UserHandler
}

func (api *API) SetupRouter() {

	// User
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())

}
