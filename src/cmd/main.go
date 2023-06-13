package main

import (
	"log"

	"github.com/krmsaeed/golang-web-api/api"
	"github.com/krmsaeed/golang-web-api/config"
	"github.com/krmsaeed/golang-web-api/data/cache"
	"github.com/krmsaeed/golang-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
