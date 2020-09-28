package db

import (
	"ecommerce-backend/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db *sqlx.DB
}

func (s *Sql) Connect(cfg *model.Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbUserName,
		cfg.Database.DbPassword,
		cfg.Database.DbName)

	s.Db = sqlx.MustConnect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println("connect success")
}

func (s *Sql) Close() {
	s.Db.Close()
}
