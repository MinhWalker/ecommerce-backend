package router

import (
	"ecommerce-backend/handler"
	"ecommerce-backend/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	AdminHandler handler.AdminHandler
}

func (api *API) SetupRouter() {
	// User
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())

}

func (api *API) SetupAdminRouter() {
	// Admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.AdminHandler.GenToken)
	admin.POST("/sign-up", api.AdminHandler.HandleSignUp)
	admin.POST("/sign-in", api.AdminHandler.HandleSignIn)

}
