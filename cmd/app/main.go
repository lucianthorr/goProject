package main

import (
	"log"

	"github.com/lucianthorr/goProject/api"
	"github.com/lucianthorr/goProject/configs"
	"github.com/lucianthorr/goProject/db"
)

func main() {
	cfg, err := configs.Read("configs/dev.yaml")
	if err != nil {
		log.Fatal(err)
	}
	dbCli := db.New(cfg.DB)

	api := api.New(cfg.API, dbCli)

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
