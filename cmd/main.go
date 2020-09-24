package main

import (
	"ecommerce-backend/db"
	"ecommerce-backend/model"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func main()  {
	var cfg model.Config
	loadConfig(&cfg)

	var sql = new(db.Sql)
	sql.Connect(&cfg)
	defer sql.Close()

}

func loadConfig(cfg *model.Config) {
	f, err := os.Open("../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
