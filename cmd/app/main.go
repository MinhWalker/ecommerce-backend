package main

import (
	"ecommerce-backend/db"
	"ecommerce-backend/handler"
	"ecommerce-backend/model"
	"ecommerce-backend/repository"
	"ecommerce-backend/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

func main() {
	var cfg model.Config
	loadConfig(&cfg)
	setupEnv(&cfg)

	var sql = new(db.Sql)
	sql.Connect(&cfg)
	defer sql.Close()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// User and Admin
	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	adminHandler := handler.AdminHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	// Product
	cateHandler := handler.CateHandler{
		CateRepo: repository.NewCateRepo(sql),
	}

	productHandler := handler.ProductHandler{
		ProductRepo: repository.NewProductRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		AdminHandler: adminHandler,
		CateHandler: cateHandler,
		ProductHandler: productHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}

func setupEnv(cfg *model.Config) {
	jwtExpires := strconv.Itoa(cfg.Server.JwtExpires)
	os.Setenv("JwtExpires", jwtExpires)
}

func loadConfig(cfg *model.Config) {
	f, err := os.Open("../../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
