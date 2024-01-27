package cmd

import (
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/config"
	"github.com/Raajheer1/tamuhack-x/m/v2/pkg/database"
)

func main() {
	cfg := config.New()

	db := database.Connect(cfg.Database)
	models.AutoMigrate(db)
}
