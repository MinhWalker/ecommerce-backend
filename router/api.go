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
	CateHandler handler.CateHandler
}

func (api *API) SetupRouter() {
	// User
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())
	user.PUT("/update", api.UserHandler.HandleUpdateUsers, middleware.JWTMiddleware())

	// categories
	categories := api.Echo.Group("/cate", middleware.JWTMiddleware(), middleware.CheckAdminRole())
	categories.POST("/add", api.CateHandler.HandleAddCate)
	categories.PUT("/edit", api.CateHandler.HandleEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandleCateDetail)
	categories.GET("/list", api.CateHandler.HandleCateList)
	categories.DELETE("/delete", api.CateHandler.HandleDeleteCate)

	// product
	//product := api.Echo.Group("/product",
	//	middleware.CheckAdminRole(),
	//	middleware.JWTMiddleware())

}

func (api *API) SetupAdminRouter() {
	// Admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.AdminHandler.GenToken)
	admin.POST("/sign-up", api.AdminHandler.HandleSignUp, middleware.JWTMiddleware())
	admin.POST("/sign-in", api.AdminHandler.HandleSignIn)

}
