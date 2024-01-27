package main

import (
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/config"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	"log"
)

func main() {
	cfg := config.New()

	db := database.Connect(cfg.Database)
	err := models.AutoMigrate(db)
	if err != nil {
		log.Fatalf("[Database] Migration Error: %s", err)
	}

}
