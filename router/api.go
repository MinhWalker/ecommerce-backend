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
	ProductHandler handler.ProductHandler
}

func (api *API) SetupRouter() {
	// User
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())
	user.PUT("/update", api.UserHandler.HandleUpdateUsers, middleware.JWTMiddleware())
	user.DELETE("/delete/:id", api.UserHandler.HandleDeleteUser, middleware.JWTMiddleware())

	// categories
	categories := api.Echo.Group("/cate", middleware.JWTMiddleware(), middleware.CheckAdminRole())
	categories.POST("/add", api.CateHandler.HandleAddCate)
	categories.PUT("/edit", api.CateHandler.HandleEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandleCateDetail)
	categories.GET("/list", api.CateHandler.HandleCateList)
	categories.DELETE("/delete/:id", api.CateHandler.HandleDeleteCate)		//update last

	// product
	product := api.Echo.Group("/product", middleware.JWTMiddleware(), middleware.CheckAdminRole())
	product.POST("/add", api.ProductHandler.HandleAddProduct)
	product.GET("/detail/:id", api.ProductHandler.HandleProductDetail)
	product.GET("/list", api.ProductHandler.HandleProductList)
	product.PUT("/edit", api.ProductHandler.HandleEditProduct)
	product.DELETE("/delete/:id", api.ProductHandler.HandleDeleteProduct)

}

func (api *API) SetupAdminRouter() {
	// Admin
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.AdminHandler.GenToken)
	admin.POST("/sign-up", api.AdminHandler.HandleSignUp, middleware.JWTMiddleware())
	admin.POST("/sign-in", api.AdminHandler.HandleSignIn)
	//admin.DELETE("/delete/:id", api.AdminHandler)

}
