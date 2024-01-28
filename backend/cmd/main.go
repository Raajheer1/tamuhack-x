package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/config"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database/models"
	v1 "github.com/Raajheer1/tamuhack-x/m/v2/v1"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "54321!")
	os.Setenv("DB_DATABASE", "tamuhack")

	cfg := config.New()

	database.DB = database.Connect(cfg.Database)
	err := models.AutoMigrate(database.DB)
	if err != nil {
		log.Fatalf("[Database] Migration Error: %s", err)
	}

	s := v1.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s))

	// acc := models.Account{ID: "12",
	// 	FirstName: "Johny",
	// 	LastName:  "Test",
	// 	Email:     "jt@gmail.com",
	// 	Money:     423}

	// acc.Create(database.DB)
}
