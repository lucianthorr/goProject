package main

import (
	"log"

	"github.com/lucianthorr/goProject/api"
	"github.com/lucianthorr/goProject/db"
)

func main() {
	dbCfg := &db.Config{}
	dbCli := db.New(dbCfg)

	apiCfg := &api.Config{
		dbCli: dbCli,
	}
	api := api.New(apiCfg)
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
