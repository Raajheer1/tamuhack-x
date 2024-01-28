package main

import (
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/config"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	v1 "github.com/Raajheer1/tamuhack-x/m/v2/v1"
	"log"
	"net/http"
)

func main() {

	cfg := config.New()

	database.DB = database.Connect(cfg.Database)
	err := models.AutoMigrate(database.DB)
	if err != nil {
		log.Fatalf("[Database] Migration Error: %s", err)
	}

	s := v1.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s))
}
